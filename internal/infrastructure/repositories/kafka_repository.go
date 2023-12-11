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

func (k *KafkaRepository) Post(message string) bool {
	write, err := k.connection.Write([]byte(message))
	if err != nil {
		panic(err)
	}

	return write > 0
}
