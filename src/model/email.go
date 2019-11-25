package model

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/asaskevich/govalidator"

	"github.com/xdhuxc/xdhuxc-message/src/utils"
)

type Email struct {
	Sender      string     `json:"sender"`
	User        string     `json:"-"`
	Receivers   utils.JSON `json:"receivers"`
	Cc          utils.JSON `json:"cc"`
	ContentType string     `json:"contentType"`
	Subject     string     `json:"subject"`
	Content     string     `json:"content"`
	Description string     `json:"description"`
	CreateTime  time.Time  `json:"create_time"`
	UpdateTime  time.Time  `json:"update_time"`
}

func (e *Email) Validate() error {
	if r := govalidator.Trim(e.Sender, " "); govalidator.IsNull(r) {
		return errors.New("the user of email must be specified")
	}

	if len(e.Receivers) < 1 {
		return errors.New("the receiver of email must be specified at least one")
	}

	if r := govalidator.Trim(e.ContentType, " "); govalidator.IsNull(r) {
		return errors.New("the user of email must be specified")
	}

	if r := govalidator.Trim(e.Subject, " "); govalidator.IsNull(r) {
		return errors.New("the subject of email must be specified")
	}

	if r := govalidator.Trim(e.Content, " "); govalidator.IsNull(r) {
		return errors.New("the body of email must be specified")
	}

	return nil
}

func (e *Email) String() string {
	if dataInBytes, err := json.Marshal(&e); err == nil {
		return string(dataInBytes)
	}

	return ""
}

func (e *Email) Convert() *Message {
	return &Message{
		Sender:      e.Sender,
		User:        e.User,
		Receivers:   e.Receivers,
		Cc:          e.Receivers,
		EmailType:   e.ContentType,
		Subject:     e.Subject,
		Content:     e.Content,
		Description: e.Description,
	}
}
