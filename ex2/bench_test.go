package main

import (
	context "context"
	"testing"
)

func BenchmarkFoo(b *testing.B) {
	setup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		foo()
	}
	b.StopTimer()
	cleanup()
}

func BenchmarkSayHello(b *testing.B) {
	s := &server{}
	for i := 0; i < b.N; i++ {
		s.SayHello(context.Background(), &HelloRequest{
			Name: "Aurélien DEROIDE",
		})
	}
}
