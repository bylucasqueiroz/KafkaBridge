package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/bylucasqueiroz/kafka-bridge/mocks"
	"github.com/bylucasqueiroz/kafka-bridge/service"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestProcessorSuccess(t *testing.T) {
	consumer := &mocks.MockConsumer{Messages: []string{"msg1", "msg2"}}
	producer := &mocks.MockProducer{}
	logger := zap.NewNop()

	proc := service.NewProcessor(consumer, producer, logger)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	proc.Start(ctx, 1)

	assert.Equal(t, []string{"msg1", "msg2"}, producer.Produced)
}

func TestProcessorConsumeError(t *testing.T) {
	consumer := &mocks.MockConsumer{Fail: true}
	producer := &mocks.MockProducer{}
	logger := zap.NewNop()

	proc := service.NewProcessor(consumer, producer, logger)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	proc.Start(ctx, 1)
	assert.Empty(t, producer.Produced)
}

func TestProcessorProduceError(t *testing.T) {
	consumer := &mocks.MockConsumer{Messages: []string{"msg1"}}
	producer := &mocks.MockProducer{Fail: true}
	logger := zap.NewNop()

	proc := service.NewProcessor(consumer, producer, logger)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	proc.Start(ctx, 1)
	assert.Empty(t, producer.Produced)
}
