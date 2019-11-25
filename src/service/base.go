package service

import (
	"github.com/jinzhu/gorm"
	"github.com/xdhuxc/xdhuxc-message/src/model"
)

type BaseService struct {
	AuditService       *auditService
	MessageService     *messageService
	EmailService       *emailService
	DingTalkService    *dingTalkService
	HealthCheckService *healthCheckService
}

func NewBaseService(dta model.DingTalkAuthentication, emailServer model.EmailServer, db *gorm.DB) *BaseService {
	return &BaseService{
		AuditService:       newAuditService(db),
		MessageService:     newMessageService(db),
		EmailService:       newEmailService(emailServer),
		DingTalkService:    newDingTalkService(dta),
		HealthCheckService: newHelathCheckService(db),
	}
}
