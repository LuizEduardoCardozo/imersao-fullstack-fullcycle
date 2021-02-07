package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/application/grpc/pb"
	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/application/usecase"
	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/infra/repository"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// StartGrpcServer starts the grpc server into a specific port
func StartGrpcServer(database *gorm.DB, port int) {

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repository.PixKeyRepositoryDB{DB: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)

	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

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
