package mail

// Client the mail client interface
type Client interface {
	// Send Send a mail message
	Send(m Message) error
}
