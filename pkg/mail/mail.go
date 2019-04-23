package mail

import (
	"crypto/tls"
	"net/smtp"
)

// NewClient create a new client
func NewClient(c Config) Client {
	return &client{config: c}
}

func (c *client) Send(m Message) error {

	m.trim()

	client, err := smtpClient(c.config)
	if err != nil {
		return err
	}

	defer client.Close()

	if c.config.WithAuth() {

		auth := smtp.PlainAuth("", c.config.AuthUser(m.From), c.config.Password, c.config.Host)
		if err := client.Auth(auth); err != nil {
			return err
		}
	}

	if err := client.Mail(m.From); err != nil {
		return err
	}

	if err := addRcpt(client, m.To); err != nil {
		return err
	}
	if err := addRcpt(client, m.CC); err != nil {
		return err
	}
	if err := addRcpt(client, m.BCC); err != nil {
		return err
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

func addRcpt(c *smtp.Client, a []string) error {
	for _, addr := range a {
		if err := c.Rcpt(addr); err != nil {
			return err
		}
	}
	return nil
}

func smtpClient(c Config) (*smtp.Client, error) {
	if c.NoTLS {
		return smtp.Dial(c.ServerName())
	}

	client, err := smtp.Dial(c.ServerName())
	if err != nil {
		return nil, err
	}

	tlsconfig := &tls.Config{
		InsecureSkipVerify: c.SkipTLSVerify,
		ServerName:         c.Host,
	}

	err = client.StartTLS(tlsconfig)

	if err != nil {
		return nil, err
	}

	return client, nil
}
