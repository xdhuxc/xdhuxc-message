package model

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/asaskevich/govalidator"

	"github.com/xdhuxc/xdhuxc-message/src/utils"
)

type DingTalk struct {
	Sender      string     `json:"sender"`
	User        string     `json:"-"`
	Receivers   utils.JSON `json:"receivers"`
	Message     string     `json:"message"`
	Description string     `json:"description"`
	CreateTime  time.Time  `json:"create_time"`
	UpdateTime  time.Time  `json:"update_time"`
}

func (dt *DingTalk) Validate() error {
	if r := govalidator.Trim(dt.Sender, " "); govalidator.IsNull(r) {
		return errors.New("the user of dingtalk message must be specified")
	}

	if len(dt.Receivers) < 1 {
		return errors.New("the user of dingtalk message must be specified")
	}

	if r := govalidator.Trim(dt.Message, " "); govalidator.IsNull(r) {
		return errors.New("the user of dingtalk message must be specified")
	}

	return nil
}

func (dt *DingTalk) String() string {
	if dataInBytes, err := json.Marshal(&dt); err == nil {
		return string(dataInBytes)
	}

	return ""
}

func (dt *DingTalk) Convert() *Message {
	return &Message{
		Sender:      dt.Sender,
		User:        dt.User,
		Receivers:   dt.Receivers,
		Content:     dt.Message,
		Description: dt.Description,
	}
}

// 发送钉钉消息辅助结构体
type DingTalkMessage struct {
	AgentID             int64               `json:"agent_id"`
	UserIDList          string              `json:"userid_list"`
	DepartmentIDList    string              `json:"dept_id_list,omitempty"`
	ToAllUser           bool                `json:"to_all_user,omitempty"`
	DingTalkMessageBody DingtalkMessageBody `json:"msg"`
}

type DingtalkMessageBody struct {
	DingTalkMessageType string              `json:"msgtype"`
	DingTalkMessageText DingTalkMessageText `json:"text"`
}

type DingTalkMessageText struct {
	Content string `json:"content"`
}

func (dtm *DingTalkMessage) String() string {
	if dataInBytes, err := json.Marshal(&dtm); err == nil {
		return string(dataInBytes)
	}

	return ""
}
