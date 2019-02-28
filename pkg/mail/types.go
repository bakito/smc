package mail

import (
	"bytes"
	"fmt"
	"strings"
)

type client struct {
	config Config
}

type Config struct {
	Host     string
	Port     uint
	User     string
	Password string
}

func (c *Config) ServerName() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Config) WithAuth() bool {
	return c.Password != ""
}

func (c *Config) AuthUser(from string) string {
	if c.User != "" {
		return c.User
	}
	return from
}

type Message struct {
	From        string
	To          []string
	CC          []string
	BCC         []string
	Subject     string
	Body        string
	ContentType string
	Encoding    string
}

func (m *Message) trim() {
	m.From = strings.TrimSpace(m.From)
	m.Subject = strings.TrimSpace(m.Subject)
	m.ContentType = strings.TrimSpace(m.ContentType)
	m.Encoding = strings.TrimSpace(m.Encoding)
	trim(m.To)
	trim(m.CC)
	trim(m.BCC)
}

func (m *Message) message() []byte {

	var b bytes.Buffer

	b.WriteString("From: ")
	b.WriteString(m.From)
	b.WriteString("\r\n")

	b.WriteString("To: ")
	b.WriteString(strings.Join(m.To, ","))
	b.WriteString("\r\n")

	b.WriteString("Subject: ")
	b.WriteString(m.Subject)
	b.WriteString("\r\n")

	b.WriteString("MIME-version: 1.0\r\n")
	b.WriteString("Content-Type: ")
	b.WriteString(m.ContentType)
	b.WriteString("; charset=\"")
	b.WriteString(m.Encoding)
	b.WriteString("\"\r\n\r\n")

	b.WriteString(m.Body)
	b.WriteString("\r\n")
	return b.Bytes()
}

func trim(sl []string) {
	for i, s := range sl {
		sl[i] = strings.TrimSpace(s)
	}
}
