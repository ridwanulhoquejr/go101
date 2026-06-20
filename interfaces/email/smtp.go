package email

import (
	"fmt"
	"net/smtp"

	"example.com/go101/interfaces/notification"
)

type SMTPClient struct {
	host string
	from string
	port string
	auth smtp.Auth
}

func NewSMTPClient(host, from, port string) *SMTPClient {
	auth := smtp.PlainAuth("", "dummy", "12345", host)
	return &SMTPClient{host: host, from: from, port: port, auth: auth}
}

// interface compilance
var _ notification.EmailDispatcher = (*SMTPClient)(nil)

func (s *SMTPClient) Send(to, subject, body string) error {

	msg := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))
	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	return smtp.SendMail(addr, s.auth, s.from, []string{to}, msg)
}
