package mocks

import "errors"

type MockProducer struct {
    Produced []string
    Fail     bool
}

var ErrMockProduce = errors.New("mock produce error")

func (m *MockProducer) Produce(message string) error {
    if m.Fail {
        return ErrMockProduce
    }
    m.Produced = append(m.Produced, message)
    return nil
}

func (m *MockProducer) Close() {}
