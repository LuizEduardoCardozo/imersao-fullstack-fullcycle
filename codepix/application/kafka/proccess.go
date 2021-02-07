package kafka

import (
	"fmt"

	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/application/factory"
	appmodel "github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/application/model"
	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/application/usecase"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
)

// KafkaProcessor helper struct
type KafkaProcessor struct {
	Database     *gorm.DB
	Producer     *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

// NewKafkaProcessor creates a new Kafka Processor
func NewKafkaProcessor(database *gorm.DB, producer *ckafka.Producer, deliveryChan chan ckafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		Database:     database,
		Producer:     producer,
		DeliveryChan: deliveryChan,
	}
}

// Consume an message From Kafka
func (k *KafkaProcessor) Consume() error {

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "consumergroup",
		"auto.offset.reset": "earliest",
	}

	c, err := ckafka.NewConsumer(configMap)

	if err != nil {
		return err
	}

	topics := []string{"transactions", "transaction_confirmation"}

	c.SubscribeTopics(topics, nil)

	fmt.Println("Kafka consumer has been started!")

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			k.processMessage(msg)
			//fmt.Println(string(msg.Value)) // To print the message
		}

	}

}

// processMessage process the kafka message passed through arguments
func (k *KafkaProcessor) processMessage(msg *ckafka.Message) {

	var (
		transactionTopics             string = "transactions"
		transactionConfirmationTopics string = "transaction_confirmation"
	)

	switch topic := *msg.TopicPartition.Topic; topic {
	case transactionTopics:
		k.processTransation(msg)
	case transactionConfirmationTopics:
		k.processTransationConfirmation(msg)
	default:
		fmt.Println("is not a valid topic", string(msg.Value))
	}

}

func (k *KafkaProcessor) processTransation(msg *ckafka.Message) error {

	transaction := appmodel.NewTransaction()
	err := transaction.ParseJSON(msg.Value)

	if err != nil {
		return err
	}

	transactionUseCase := factory.TransactionUseCaseFactory(k.Database)

	createdTransaction, err := transactionUseCase.Register(
		transaction.AccountID,
		transaction.Amount,
		transaction.PixKeyTo,
		transaction.PixKeyKindTo,
		transaction.Description,
	)

	if err != nil {
		fmt.Println("Error while registering a transaction", err)
		return err
	}

	topic := "bank" + createdTransaction.PixKeyTo.Account.Bank.Code
	transaction.ID = createdTransaction.ID
	transaction.Status = "pending"

	transactionJSON, err := transaction.ToJSON()

	if err != nil {
		return err
	}

	err = Publish(string(transactionJSON), topic, k.Producer, k.DeliveryChan)

	if err != nil {
		return err
	}

	return nil

}

// processTransationConfirmation the kafka message
func (k *KafkaProcessor) processTransationConfirmation(msg *ckafka.Message) error {

	transaction := appmodel.NewTransaction()
	err := transaction.ParseJSON(msg.Value)

	if err != nil {
		return err
	}

	transactionUseCase := factory.TransactionUseCaseFactory(k.Database)

	if transaction.Status == "confirmed" {

		err = k.confirmTransaction(transaction, transactionUseCase)

		if err != nil {
			return err
		}

	} else if transaction.Status == "completed" {

		_, err = transactionUseCase.Completed(transaction.ID)

		if err != nil {
			return err
		}

		return nil

	}

	return nil

}

func (k *KafkaProcessor) confirmTransaction(transaction *appmodel.Transaction, transactionUseCase usecase.TransactionUseCase) error {

	confirmedTransaction, err := transactionUseCase.Confirm(transaction.ID)
	if err != nil {
		return nil
	}

	topic := "bank" + confirmedTransaction.AccountFrom.Bank.Code
	transactionJSON, err := transaction.ToJSON()

	if err != nil {
		return err
	}

	err = Publish(string(transactionJSON), topic, k.Producer, k.DeliveryChan)

	if err != nil {
		return err
	}

	return nil

}
