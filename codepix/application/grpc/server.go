package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

// StartGrpcServer starts the grpc server into a specific port
func StartGrpcServer(database *gorm.DB, port int) {

	grpcServer := grpc.NewServer()

	address := fmt.Sprintf("0.0.0.0:%d", port)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("Cannot start the server - ", err)
	}

	log.Printf("gRPC server running on port %d", port)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot start gRPC server - ", err)
	}

}
