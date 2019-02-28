package mail

import (
	"bytes"
	"fmt"
	"strings"
)

type client struct {
	config Config
}

// Config mail client config
type Config struct {
	Host          string
	Port          uint
	User          string
	Password      string
	SkipTLSVerify bool
}

// ServerName get host and port
func (c *Config) ServerName() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// WithAuth true, if password is set
func (c *Config) WithAuth() bool {
	return c.Password != ""
}

// AuthUser if user ist set return the user, else return the provided from
func (c *Config) AuthUser(from string) string {
	if c.User != "" {
		return c.User
	}
	return from
}

// Message an email message content
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
	m.To = splitAndTrim(m.To)
	m.CC = splitAndTrim(m.CC)
	m.BCC = splitAndTrim(m.BCC)
}

func (m *Message) message() []byte {

	var b bytes.Buffer

	b.WriteString("From: ")
	b.WriteString(m.From)
	b.WriteString("\r\n")

	b.WriteString("To: ")
	b.WriteString(strings.Join(m.To, ","))
	b.WriteString("\r\n")

	if len(m.CC) > 0 {
		b.WriteString("CC: ")
		b.WriteString(strings.Join(m.CC, ","))
		b.WriteString("\r\n")
	}

	// BCC is not added to the headers

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

func splitAndTrim(in []string) []string {
	out := []string{}
	for _, s := range in {
		for _, part := range strings.Split(s, " ") {
			trimmed := strings.TrimSpace(part)
			if trimmed != "" {
				out = append(out, trimmed)
			}
		}
	}
	return out
}
