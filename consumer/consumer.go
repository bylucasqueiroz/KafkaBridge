package consumer

import (
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
    consumer *kafka.Consumer
}

func NewKafkaConsumer(brokers, groupID, topic string) (*KafkaConsumer, error) {
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": brokers,
        "group.id":          groupID,
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        return nil, err
    }

    if err := c.SubscribeTopics([]string{topic}, nil); err != nil {
        return nil, err
    }

    return &KafkaConsumer{consumer: c}, nil
}

func (kc *KafkaConsumer) Consume() (string, error) {
    msg, err := kc.consumer.ReadMessage(-1)
    if err != nil {
        return "", err
    }
    return string(msg.Value), nil
}

func (kc *KafkaConsumer) Close() {
    kc.consumer.Close()
}
