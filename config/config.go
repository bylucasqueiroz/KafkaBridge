package config

import (
    "log"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type KafkaConfig struct {
    Brokers     string
    GroupID     string
    InputTopic  string
    OutputTopic string
    NumWorkers  int
}

func LoadConfig() *KafkaConfig {
    if _, err := os.Stat(".env.local"); err == nil {
        if err := godotenv.Load(".env.local"); err != nil {
            log.Fatalf("Error loading .env.local: %v", err)
        }
    }

    brokers := os.Getenv("KAFKA_BROKERS")
    groupID := os.Getenv("KAFKA_GROUP_ID")
    inputTopic := os.Getenv("KAFKA_INPUT_TOPIC")
    outputTopic := os.Getenv("KAFKA_OUTPUT_TOPIC")
    numWorkersStr := os.Getenv("NUM_WORKERS")

    numWorkers, err := strconv.Atoi(numWorkersStr)
    if err != nil || numWorkers <= 0 {
        numWorkers = 1
    }

    return &KafkaConfig{
        Brokers:     brokers,
        GroupID:     groupID,
        InputTopic:  inputTopic,
        OutputTopic: outputTopic,
        NumWorkers:  numWorkers,
    }
}
