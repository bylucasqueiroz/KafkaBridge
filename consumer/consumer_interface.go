package consumer

type Consumer interface {
    Consume() (string, error)
    Close()
}
