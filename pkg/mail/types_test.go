package mail

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Config_ServerName(t *testing.T) {
	c := Config{
		Host: "server", Port: 123,
	}
	assert.Equal(t, "server:123", c.ServerName())
}

func Test_Message_trim(t *testing.T) {
	m := Message{
		From:        " from  \t",
		To:          []string{"  aaa bbb ccc"},
		CC:          []string{" \t 123 \n 456 890"},
		BCC:         nil,
		Subject:     " subject \n",
		ContentType: " contentType ",
		Encoding:    "\tencoding",
	}
	m.trim()

	assert.Equal(t, "from", m.From)

	assert.Len(t, m.To, 3)
	assert.Equal(t, "aaa", m.To[0])
	assert.Equal(t, "bbb", m.To[1])
	assert.Equal(t, "ccc", m.To[2])

	assert.Len(t, m.CC, 3)
	assert.Equal(t, "123", m.CC[0])
	assert.Equal(t, "456", m.CC[1])
	assert.Equal(t, "890", m.CC[2])

	assert.Len(t, m.BCC, 0)

	assert.Equal(t, "subject", m.Subject)
	assert.Equal(t, "contentType", m.ContentType)
	assert.Equal(t, "encoding", m.Encoding)
}

func Test_Message_message_max(t *testing.T) {
	m := Message{
		From:        "from",
		To:          []string{"to1", "to2"},
		CC:          []string{"cc"},
		BCC:         []string{"bcc"},
		Subject:     " subject",
		Body:        "body",
		ContentType: "contentType",
		Encoding:    "encoding",
	}

	text := string(m.message())
	lines := strings.Split(text, "\r\n")

	assert.Len(t, lines, 9)
	assert.Equal(t, "From: from", lines[0])
	assert.Equal(t, "To: to1,to2", lines[1])
	assert.Equal(t, "CC: cc", lines[2])
	assert.Equal(t, "Subject:  subject", lines[3])
	assert.Equal(t, "MIME-version: 1.0", lines[4])
	assert.Equal(t, "Content-Type: contentType; charset=\"encoding\"", lines[5])
	assert.Equal(t, "", lines[6])
	assert.Equal(t, "body", lines[7])
	assert.Equal(t, "", lines[8])
}

func Test_Message_message_min(t *testing.T) {
	m := Message{
		From:        "from",
		To:          []string{"to1", "to2"},
		Subject:     " subject",
		Body:        "body",
		ContentType: "contentType",
		Encoding:    "encoding",
	}

	text := string(m.message())
	lines := strings.Split(text, "\r\n")

	assert.Len(t, lines, 8)
	assert.Equal(t, "From: from", lines[0])
	assert.Equal(t, "To: to1,to2", lines[1])
	assert.Equal(t, "Subject:  subject", lines[2])
	assert.Equal(t, "MIME-version: 1.0", lines[3])
	assert.Equal(t, "Content-Type: contentType; charset=\"encoding\"", lines[4])
	assert.Equal(t, "", lines[5])
	assert.Equal(t, "body", lines[6])
	assert.Equal(t, "", lines[7])
}
