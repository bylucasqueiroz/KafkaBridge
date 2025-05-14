package mocks

import (
	"errors"
	"time"
)

type MockConsumer struct {
	Messages []string
	Index    int
	Fail     bool
}

var ErrMockConsume = errors.New("mock consume error")

func (m *MockConsumer) Consume() (string, error) {
	if m.Fail {
		return "", ErrMockConsume
	}
	if m.Index >= len(m.Messages) {
		time.Sleep(10 * time.Millisecond) // Evita CPU 100%
		return "", nil                    // ou return "", io.EOF se preferir
	}
	msg := m.Messages[m.Index]
	m.Index++
	return msg, nil
}

func (m *MockConsumer) Close() {}
