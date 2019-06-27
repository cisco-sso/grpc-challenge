package main

import (
	"context"
	"log"
	"net"

	"github.com/cisco-sso/grpc-challenge/pkg/generated"
	"google.golang.org/grpc"
)

type reverseServer struct{}

func main() {
	srv := grpc.NewServer()
	var server reverseServer
	generated.RegisterStringmanipServer(srv, server)
	l, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("could not listen to port :50001: %v", err)
	}
	log.Fatal(srv.Serve(l))
}

// Reverse takes a string and returns the reverse
//ex: "One small step for man" -> "nam rof pets llams enO"
func (reverseServer) Reversed(ctx context.Context, s *generated.StringRequest) (*generated.StringResponse, error) {
	a := []rune(s.Content)
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	resp := new(generated.StringResponse)
	resp.Content = string(a)
	return resp, nil
}
