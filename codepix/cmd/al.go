package cmd

import (
	"os"

	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/application/grpc"
	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/application/kafka"
	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/infra/db"
	"github.com/spf13/cobra"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

var grpcPortNumber int

// alCmd represents the al command
var alCmd = &cobra.Command{

	Use:   "al",
	Short: "Run gRPC and Apache Kafka Consumer",

	Run: func(cmd *cobra.Command, args []string) {

		database := db.ConnectDB(os.Getenv("dev"))
		grpc.StartGrpcServer(database, grpcPortNumber)

		deliveryChannel := make(chan ckafka.Event)
		producer, _ := kafka.NewKafkaProducer()

		//kafka.Publish("Hwllo, world!", "teste", producer, deliveryChannel) // Test message
		go kafka.DeliveryReport(deliveryChannel)

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChannel)
		kafkaProcessor.Consume()

	},
}

func init() {
	rootCmd.AddCommand(alCmd)
	alCmd.Flags().IntVarP(&grpcPortNumber, "gPort", "p", 50051, "set port of gRPC server")

}
