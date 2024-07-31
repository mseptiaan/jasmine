package main

import (
	"flag"
	"github.com/mseptiaan/jasmine/internal/gateway"
	"log"
)

var serverAddress = flag.String(
	"server-address",
	"dns:///0.0.0.0:10000",
	"The address to the gRPC server, in the gRPC standard naming format. "+
		"See https://github.com/grpc/grpc/blob/master/doc/naming.md for more information.",
)

func main() {
	flag.Parse()

	err := gateway.Run(*serverAddress)
	log.Fatalln(err)
}
