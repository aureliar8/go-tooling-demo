package main

import (
	context "context"
	"fmt"
	"log"
	"net"
	"net/http"

	_ "net/http/pprof"

	grpc "google.golang.org/grpc"
)

// server is used to implement GreeterServer.
type server struct {
	UnimplementedGreeterServer
}

// SayHello implements GreeterServer
func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	hash := nameHash(in.Name)
	score := nameScore(in.Name)
	reply := formatReply(in.Name, hash, score)
	return &HelloReply{
		Message: reply,
	}, nil
}

func main() {
	go func() {
		log.Printf("pprof endpoint listening at localhost:8080/debug/pprof/")
		http.ListenAndServe("localhost:8080", nil)
	}()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1234))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
