package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bylucasqueiroz/kafka-bridge/config"
	"github.com/bylucasqueiroz/kafka-bridge/consumer"
	"github.com/bylucasqueiroz/kafka-bridge/producer"
	"github.com/bylucasqueiroz/kafka-bridge/service"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg := config.LoadConfig()
	logger.Info("Starting Kafka processor",
		zap.String("brokers", cfg.Brokers),
		zap.String("group", cfg.GroupID),
		zap.String("input_topic", cfg.InputTopic),
		zap.String("output_topic", cfg.OutputTopic),
		zap.Int("workers", cfg.NumWorkers),
	)

	cons, err := consumer.NewKafkaConsumer(cfg.Brokers, cfg.GroupID, cfg.InputTopic)
	if err != nil {
		logger.Fatal("Failed to init consumer", zap.Error(err))
	}
	defer cons.Close()

	prod, err := producer.NewKafkaProducer(cfg.Brokers, cfg.OutputTopic)
	if err != nil {
		logger.Fatal("Failed to init producer", zap.Error(err))
	}
	defer prod.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	p := service.NewProcessor(cons, prod, logger)

	var wg sync.WaitGroup
	for i := 0; i < cfg.NumWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			p.Start(ctx, id)
		}(i)
	}

	<-stop
	logger.Info("Shutdown signal received. Stopping workers...")
	cancel()
	wg.Wait()
	logger.Info("All workers stopped. Exiting.")
}
