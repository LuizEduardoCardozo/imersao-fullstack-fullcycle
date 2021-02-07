/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/application/grpc"
	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/infra/db"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

var portNumber int
var database *gorm.DB

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
