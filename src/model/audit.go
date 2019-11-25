package model

import (
	"time"
)

const (
	OperatingTypeAdd    = "add"
	OperatingTypeUpdate = "update"
	OperatingTypeDelete = "delete"
	OperatingTypeLogin  = "login"
	OperatingTypeLogout = "logout"
)

const (
	MessageTypeEmail    = "email"
	MessageTypeDingTalk = "dingtalk"
	MessageTypeWeChat   = "wechat"
)

type OperationAudit struct {
	ID         int64     `json:"id" gorm:"id"`
	User       string    `json:"user" gorm:"user"`
	Operate    string    `json:"operate" gorm:"operate"`
	Object     string    `json:"object" gorm:"object"`
	CreateTime time.Time `json:"createTime" gorm:"create_time"`
}

func (oa *OperationAudit) TableName() string {
	return "xdhuxc-message_audit"
}

func (oa *OperationAudit) GetOperationResource() string {
	switch oa.Object {
	case MessageTypeEmail:
		return "邮件"
	case MessageTypeDingTalk:
		return "钉钉"
	case MessageTypeWeChat:
		return "微信"
	}

	return ""
}

func (oa *OperationAudit) GetOperationType() string {
	switch oa.Operate {
	case OperatingTypeAdd:
		return "添加"
	case OperatingTypeUpdate:
		return "更新"
	case OperatingTypeDelete:
		return "删除"
	case OperatingTypeLogin:
		return "登录"
	case OperatingTypeLogout:
		return "登出"
	}

	return ""
}
