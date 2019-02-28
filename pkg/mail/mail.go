package mail

import (
	"crypto/tls"
	"net/smtp"
)

func NewClient(c Config) Client {
	return &client{config: c}
}

func (c *client) Send(m Message) error {

	m.trim()

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         c.config.Host,
	}

	conn, err := tls.Dial("tcp", c.config.ServerName(), tlsconfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, c.config.Host)
	if err != nil {
		return err
	}

	defer client.Close()

	if c.config.WithAuth() {

		auth := smtp.PlainAuth("", c.config.AuthUser(m.From), c.config.Password, c.config.Host)
		if err = client.Auth(auth); err != nil {
			return err
		}
	}

	if err = client.Mail(m.From); err != nil {
		return err
	}
	for _, addr := range m.To {
		if err = client.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := client.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(m.message())
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return client.Quit()
}
