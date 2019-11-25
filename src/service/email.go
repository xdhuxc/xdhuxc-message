package service

import (
	"encoding/json"

	"github.com/go-gomail/gomail"

	"github.com/xdhuxc/xdhuxc-message/src/model"
)

type emailService struct {
	emailServer model.EmailServer
}

func newEmailService(emailServer model.EmailServer) *emailService {
	return &emailService{emailServer}
}

func (es *emailService) Send(e model.Email) error {
	var receivers []string
	var cc []string

	err := json.Unmarshal(e.Receivers, &receivers)
	if err != nil {
		return err
	}
	err = json.Unmarshal(e.Cc, &cc)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", es.emailServer.User)
	m.SetHeader("To", receivers...)
	if len(cc) > 0 {
		m.SetHeader("Cc", cc...)
	}
	m.SetHeader("Subject", e.Subject)
	m.SetBody(e.ContentType, e.Content)

	d := gomail.NewDialer(es.emailServer.SMTPServer, es.emailServer.Port, es.emailServer.User, es.emailServer.Password)

	return d.DialAndSend(m)
}
