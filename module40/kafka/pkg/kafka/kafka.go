package kafka

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// Broker представляет кафку брокера.
type Broker struct {
	host        string
	port        string
	readTimeout time.Duration
}

func getKafkaReadTimeout() time.Duration {
	readTimeoutStr := os.Getenv("KAFKA_READ_TIMEOUT")

	const defaultTimeout = 10 * time.Second

	if readTimeoutStr == "" {
		fmt.Println(
			"Переменная окружения KAFKA_READ_TIMEOUT не установлена, используем значение по умолчанию:", defaultTimeout,
		)
		return defaultTimeout
	}

	value, err := strconv.Atoi(readTimeoutStr)
	if err != nil {
		fmt.Printf(
			"Некорректное значение KAFKA_READ_TIMEOUT: %s, используем значение по умолчанию: %v\n", readTimeoutStr,
			defaultTimeout,
		)
		return defaultTimeout
	}

	if value <= 0 {
		fmt.Printf(
			"Значение KAFKA_READ_TIMEOUT должно быть > 0, используем значение по умолчанию: %v\n", defaultTimeout,
		)
		return defaultTimeout
	}

	return time.Duration(value) * time.Second
}

// NewBroker создает брокера.
func NewBroker() *Broker {
	_ = godotenv.Load()

	host := os.Getenv("KAFKA_HOST")
	port := os.Getenv("KAFKA_PORT")
	readTimeout := getKafkaReadTimeout()
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "9092"
	}

	return &Broker{
		host:        host,
		port:        port,
		readTimeout: readTimeout,
	}
}

// Address возвращает адрес подключения к брокеру
func (b *Broker) Address() string {
	return b.host + ":" + b.port
}

// ListTopics возвращает список топиков.
func (b *Broker) ListTopics() ([]string, error) {
	address := b.Address()
	conn, err := kafka.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		return nil, err
	}

	topicSet := make(map[string]bool)
	for _, p := range partitions {
		if !strings.HasPrefix(p.Topic, "__") { // Исключаем системные топики
			topicSet[p.Topic] = true
		}
	}

	var topics []string
	for topic := range topicSet {
		topics = append(topics, topic)
	}

	return topics, nil
}

// CreateTopic создает топик.
func (b *Broker) CreateTopic(topic string) error {
	conn, err := kafka.Dial("tcp", b.Address())
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, fmt.Sprintf("%d", controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		return err
	}

	return nil
}

// SendMessage отправляет сообщение в топик.
func (b *Broker) SendMessage(topic, message string) error {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(b.Address()),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	err := writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: []byte(message),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (b *Broker) ReadMessage(topic string) error {
	const delimiter = "----------------------------------------\n"
	reader := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers:  []string{b.Address()},
			Topic:    topic,
			GroupID:  "consumer-group-id",
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
		},
	)
	defer reader.Close()

	fmt.Printf("\nЧтение сообщение из топика \"%s\"\n", topic)
	fmt.Println("Нажмите Ctrl+C для остановки чтения")
	fmt.Printf(delimiter)

	ctx, cancel := context.WithTimeout(context.Background(), b.readTimeout)
	defer cancel()

	messageCount := 0

	for {
		msg, err := reader.FetchMessage(ctx)
		if err != nil {
			if ctx.Err() == context.DeadlineExceeded {
				if messageCount == 0 {
					fmt.Println("Сообщения не найдены в топике")
				} else {
					fmt.Printf("\nВсего прочитано сообщений: %d\n", messageCount)
				}
				break
			}

			return fmt.Errorf("ошибка чтения сообщения: %v", err)
		}

		messageCount++
		timestamp := msg.Time.Format("2006-01-02 15:04:05")
		fmt.Printf(
			"[%d] %s | Partition: %d, Offset: %d\n",
			messageCount, timestamp, msg.Partition, msg.Offset,
		)
		fmt.Printf("Сообщение: %s\n", string(msg.Value))
		fmt.Printf(delimiter)

		if err := reader.CommitMessages(ctx, msg); err != nil {
			fmt.Printf("Ошибка подтверждения сообщения: %v\n", err)
		}
	}
	return nil
}
