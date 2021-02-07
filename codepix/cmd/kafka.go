package cmd

import (
	"fmt"
	"os"

	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/application/kafka"
	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/infra/db"
	"github.com/spf13/cobra"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{

	Use:   "kafka",
	Short: "Start Kafka server",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Produzindo um mensagem!")

		deliveryChannel := make(chan ckafka.Event)
		database := db.ConnectDB(os.Getenv("dev"))
		producer, _ := kafka.NewKafkaProducer()

		//kafka.Publish("Hwllo, world!", "teste", producer, deliveryChannel) // Test message
		go kafka.DeliveryReport(deliveryChannel)

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChannel)
		kafkaProcessor.Consume()

	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kafkaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kafkaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
