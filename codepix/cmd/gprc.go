package cmd

import (
	"os"

	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/application/grpc"
	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/infra/db"
	"github.com/spf13/cobra"
)

var portNumber int

var gprcCmd = &cobra.Command{
	Use:   "gprc",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("dev"))
		grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
	rootCmd.AddCommand(gprcCmd)
	gprcCmd.Flags().IntVarP(&portNumber, "port", "p", 50051, "gPRC port number")
}
