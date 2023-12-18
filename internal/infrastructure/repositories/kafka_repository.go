package repositories

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

const (
	tennisCourtTopic = "tennis-court"
)

type KafkaRepository struct {
	connection *kafka.Conn
	reader     *kafka.Reader
}

func NewKafkaRepository() *KafkaRepository {
	connection, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", tennisCourtTopic, 0)
	if err != nil {
		panic(err)
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"kafka:9092"},
		Topic:     tennisCourtTopic,
		Partition: 0,
	})

	return &KafkaRepository{
		connection: connection,
		reader:     reader,
	}
}

func (k *KafkaRepository) Send(message string) bool {
	bytesWritten, err := k.connection.Write([]byte(message))
	if err != nil {
		panic(err)
	}

	return bytesWritten > 0
}

func (k *KafkaRepository) Receive() []string {
	var messages []string

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"kafka:9092"},
		Topic:     tennisCourtTopic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	readDeadline, _ := context.WithDeadline(context.Background(),
		time.Now().Add(2*time.Second))

	for {
		readMessage, err := r.ReadMessage(readDeadline)
		if err != nil {
			break
		}
		messages = append(messages, string(readMessage.Value))
	}

	return messages
}

func (k *KafkaRepository) Ping() bool {
	brokers, err := k.connection.Brokers()
	if err != nil {
		panic(err)
	}

	return len(brokers) > 0
}
