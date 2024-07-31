package main

import (
	"github.com/mseptiaan/jasmine/internal/core"
	"github.com/mseptiaan/jasmine/internal/gateway"
	"github.com/mseptiaan/jasmine/internal/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()

	pb.RegisterJasmineEndpointServer(s, core.New(false))

	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}
