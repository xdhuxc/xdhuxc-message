package service

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/xdhuxc/xdhuxc-message/src/pkg"

	"github.com/xdhuxc/xdhuxc-message/src/model"
)

type metricsService struct {
	db *gorm.DB
}

func newMetricsService(db *gorm.DB) *metricsService {
	return &metricsService{db}
}

func (ms *metricsService) Count() *model.MetricsResult {
	mr := new(model.MetricsResult)

	if err := ms.db.Model(&model.Message{}).Count(&mr.TotalMessages).Error; err != nil {
		log.Errorln(err)
	}
	if err := ms.db.Model(&model.Message{}).Where("is_sent = ?", false).Count(&mr.TotalFailedMessages).Error; err != nil {
		log.Errorln(err)
	}
	if err := ms.db.Model(&model.Message{}).Where("is_sent = ?", true).Count(&mr.TotalSentMessages).Error; err != nil {
		log.Errorln(err)
	}
	if err := ms.db.Model(&model.Message{}).Where("message_type = ?", pkg.MessageTypeDingTalk).Count(&mr.DingTalkMessages).Error; err != nil {
		log.Errorln(err)
	}
	if err := ms.db.Model(&model.Message{}).Where("message_type = ? AND is_sent = ?", pkg.MessageTypeDingTalk, false).Count(&mr.DingTalkFailedMessages).Error; err != nil {
		log.Errorln(err)
	}
	if err := ms.db.Model(&model.Message{}).Where("message_type = ? AND is_sent = ?", pkg.MessageTypeDingTalk, true).Count(&mr.DingTalkSentMessages).Error; err != nil {
		log.Errorln(err)
	}
	if err := ms.db.Model(&model.Message{}).Where("message_type = ?", pkg.MessageTypeEmail).Count(&mr.EmailMessages).Error; err != nil {
		log.Errorln(err)
	}
	if err := ms.db.Model(&model.Message{}).Where("message_type = ? AND is_sent = ?", pkg.MessageTypeEmail, false).Count(&mr.EmailFailedMessages).Error; err != nil {
		log.Errorln(err)
	}
	if err := ms.db.Model(&model.Message{}).Where("message_type = ? AND is_sent = ?", pkg.MessageTypeEmail, true).Count(&mr.EmailSentMessages).Error; err != nil {
		log.Errorln(err)
	}

	return mr
}
