package producer

import (
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
    producer *kafka.Producer
    topic    string
}

func NewKafkaProducer(brokers, topic string) (*KafkaProducer, error) {
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": brokers,
    })
    if err != nil {
        return nil, err
    }

    return &KafkaProducer{producer: p, topic: topic}, nil
}

func (kp *KafkaProducer) Produce(message string) error {
    return kp.producer.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &kp.topic, Partition: kafka.PartitionAny},
        Value:          []byte(message),
    }, nil)
}

func (kp *KafkaProducer) Close() {
    kp.producer.Flush(15000)
    kp.producer.Close()
}
