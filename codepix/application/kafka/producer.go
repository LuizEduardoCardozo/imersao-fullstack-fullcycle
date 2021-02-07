package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

//NewKafkaProducer returns a new kafka producer
func NewKafkaProducer() (*ckafka.Producer, error) {
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}

	producer, err := ckafka.NewProducer(&configMap)

	if err != nil {
		return nil, err
	}

	return producer, nil

}

// Publish a message to Apache Kafka
func Publish(msg string, topic string, producer *ckafka.Producer, deliveryChannel chan ckafka.Event) error {

	message := ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic:     &topic,
			Partition: ckafka.PartitionAny,
		},
		Value: []byte(msg),
	}

	err := producer.Produce(&message, deliveryChannel)

	if err != nil {
		return err
	}

	return nil

}

// DeliveryReport sends an menssage to deliveryChannel
func DeliveryReport(deliveryChan chan ckafka.Event) {

	for e := range deliveryChan {
		switch ev := e.(type) {
		case *ckafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery Failed", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to", ev.TopicPartition)
			}
		}
	}

}
