package producer

type Producer interface {
    Produce(message string) error
    Close()
}
