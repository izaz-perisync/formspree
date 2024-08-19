package mailmon

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"
)

const (
	mime        = "MIME-Version: 1.0\n"
	contentType = "Content-Type: text/html\n"
)

type IService interface {
	SendMail(subject, body string, recipients []string) error
}

type Service struct {
	user   string
	secret string
	host   string
	port   int
}

func New(user, secret, host string, port int) *Service {
	return &Service{
		user:   user,
		secret: secret,
		host:   host,
		port:   port,
	}
}

func (s *Service) SendMail(subject, body string, recipients []string) error {

	if err := validateMailArgs(subject, body, recipients); err != nil {
		return err
	}

	from := fmt.Sprintf("From: %s\n", s.user)
	sub := fmt.Sprintf("Subject: %s\n", subject)

	msg := []byte(from + sub + mime + contentType + body)

	auth := smtp.PlainAuth("", s.user, s.secret, s.host)
	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	return smtp.SendMail(addr, auth, s.user, recipients, msg)
}

func validateMailArgs(subject, body string, recipients []string) error {
	var errs []string
	if subject == "" {
		errs = append(errs, "subject is empty")
	}
	if body == "" {
		errs = append(errs, "body is empty")
	}
	if len(recipients) < 1 {
		errs = append(errs, "len of recipients array is < 1")
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}
	return nil
}
