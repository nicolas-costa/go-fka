package repositories

import (
	"context"
	"github.com/segmentio/kafka-go"
)

const (
	tennisCourtTopic = "tennis_court"
)

type KafkaRepository struct {
	connection *kafka.Conn
}

func NewKafkaRepository() *KafkaRepository {
	connection, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", tennisCourtTopic, 0)
	if err != nil {
		panic(err)
	}

	return &KafkaRepository{
		connection: connection,
	}
}

func (k *KafkaRepository) Send(message string) bool {
	bytesWritten, err := k.connection.Write([]byte(message))
	if err != nil {
		panic(err)
	}

	return bytesWritten > 0
}

func (k *KafkaRepository) Ping() bool {
	brokers, err := k.connection.Brokers()
	if err != nil {
		panic(err)
	}

	return len(brokers) > 0
}
