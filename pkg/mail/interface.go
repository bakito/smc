package mail

type Client interface {
	Send(m Message) error
}
