package service

import (
	"context"

	"github.com/bylucasqueiroz/kafka-bridge/consumer"
	"github.com/bylucasqueiroz/kafka-bridge/producer"
	"go.uber.org/zap"
)

type Processor struct {
	Consumer consumer.Consumer
	Producer producer.Producer
	Logger   *zap.Logger
}

func NewProcessor(c consumer.Consumer, p producer.Producer, logger *zap.Logger) *Processor {
	return &Processor{Consumer: c, Producer: p, Logger: logger}
}

func (p *Processor) Start(ctx context.Context, workerID int) {
	p.Logger.Info("Worker started", zap.Int("worker_id", workerID))

	for {
		select {
		case <-ctx.Done():
			p.Logger.Info("Worker stopping", zap.Int("worker_id", workerID))
			return
		default:
			msg, err := p.Consumer.Consume()
			if err != nil {
				p.Logger.Warn("Error consuming message", zap.Error(err), zap.Int("worker_id", workerID))
				continue
			}

			if msg == "" {
				continue
			}

			err = p.Producer.Produce(msg)
			if err != nil {
				p.Logger.Error("Error producing message", zap.Error(err), zap.Int("worker_id", workerID))
			} else {
				p.Logger.Debug("Message produced", zap.String("value", msg), zap.Int("worker_id", workerID))
			}
		}
	}
}
