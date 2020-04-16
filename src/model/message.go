package model

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/asaskevich/govalidator"

	"github.com/xdhuxc/xdhuxc-message/src/pkg"
	"github.com/xdhuxc/xdhuxc-message/src/utils"
)

type Message struct {
	// common
	ID          int64      `json:"id" gorm:"id"`
	User        string     `json:"-" gorm:"user"`
	Sender      string     `json:"sender" gorm:"sender"`
	MessageType string     `json:"messageType" gorm:"message_type"`
	IsSent      bool       `json:"isSent" gorm:"is_sent"`
	Content     string     `json:"content" gorm:"content"`
	Description string     `json:"description,omitempty" gorm:"description"`
	CreateTime  time.Time  `json:"createTime" gorm:"create_time"`
	UpdateTime  time.Time  `json:"updateTime" gorm:"update_time"`
	Receivers   utils.JSON `json:"receivers" gorm:"receivers"`
	// email
	Cc        utils.JSON `json:"cc,omitempty" gorm:"cc"`
	EmailType string     `json:"emailType,omitempty" gorm:"email_type"`
	Subject   string     `json:"subject,omitempty" gorm:"subject"`
	// dingtalk

	// phone

	// wechat
}

func (m *Message) Validate() error {
	if r := govalidator.Trim(m.Sender, " "); govalidator.IsNull(r) {
		return errors.New("the user of message must be specified")
	}

	if r := govalidator.Trim(m.MessageType, " "); govalidator.IsNull(r) {
		return errors.New("the type of message must be specified")
	}

	if m.MessageType == pkg.MessageTypeEmail {
		if r := govalidator.Trim(m.EmailType, " "); govalidator.IsNull(r) {
			return errors.New("the email type of email message must be specified")
		}

		if r := govalidator.Trim(m.Subject, " "); govalidator.IsNull(r) {
			return errors.New("the subject of email message must be specified")
		}
	}

	if r := govalidator.Trim(m.Content, " "); govalidator.IsNull(r) {
		return errors.New("the user of message must be specified")
	}

	if len(m.Receivers) < 1 {
		return errors.New("the receiver of message must be specified at least one")
	}

	return nil
}

func (m *Message) TableName() string {
	return "xdhuxc-message_message"
}

func (m *Message) String() string {
	if dataInBytes, err := json.Marshal(&m); err == nil {
		return string(dataInBytes)
	}

	return ""
}

func (m *Message) ConvertToEmail() *Email {
	return &Email{
		Sender:      m.Sender,
		User:        m.User,
		Receivers:   m.Receivers,
		Cc:          m.Cc,
		ContentType: m.EmailType,
		Subject:     m.Subject,
		Content:     m.Content,
		Description: m.Description,
	}
}

func (m *Message) ConvertToDingTalk() *DingTalk {
	return &DingTalk{
		Sender:      m.Sender,
		User:        m.Sender,
		Receivers:   m.Receivers,
		Message:     m.Content,
		Description: m.Description,
	}
}
