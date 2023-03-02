package ex3

import (
	context "context"

	"github.com/exoscale/stelling/fxgrpc"
	"github.com/exoscale/stelling/fxlogging"
	"go.uber.org/fx"
	"google.golang.org/grpc"
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

func NewHelloWorldServer() *server {
	return &server{}
}

func RegisterHelloWorldServer(s *server, grpc *grpc.Server) {
	RegisterGreeterServer(grpc, s)
}

type Conf struct {
	fxlogging.Logging
	fxgrpc.Server
	//fxpprof.Pprof
}

func main() {
	conf := &Conf{
		Server: fxgrpc.Server{
			Port: 1234,
		},
		// Pprof: fxpprof.Pprof{
		// 	Enabled: true,
		// 	Port:    8080,
		// },
	}
	opts := fx.Options(
		fxlogging.Module,
		fxgrpc.ServerModule,
		fx.Supply(
			conf,
			fx.Annotate(conf, fx.As(new(fxlogging.LoggingConfig))),
			fx.Annotate(conf, fx.As(new(fxgrpc.GrpcServerConfig))),
			//fx.Annotate(conf, fx.As(new(fxpprof.Pprof))),
		),
		fx.Provide(
			NewHelloWorldServer,
		),
		fx.Invoke(
			RegisterHelloWorldServer,
			fxgrpc.StartGrpcServer,
		),
	)

	fx.New(opts).Run()
}
