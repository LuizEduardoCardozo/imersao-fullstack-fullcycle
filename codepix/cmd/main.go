package main

import (
	"os"

	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/application/grpc"
	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/infra/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {

	database = db.ConnectDB(os.Getenv("dev"))
	grpc.StartGrpcServer(database, 50051)

}
