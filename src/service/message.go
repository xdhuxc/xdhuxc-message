package service

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/xdhuxc/xdhuxc-message/src/model"
)

type messageService struct {
	db *gorm.DB
}

func newMessageService(db *gorm.DB) *messageService {
	return &messageService{db}
}

func (ms *messageService) Create(m model.Message) (*model.Message, error) {
	if err := m.Validate(); err != nil {
		return nil, err
	}

	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()

	if err := ms.db.Create(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (ms *messageService) UpdateStatus(m model.Message) error {
	if err := ms.db.Model(&m).Update(map[string]interface{}{
		"is_sent": true,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (ms *messageService) GetMessageByID(id int64) (*model.Message, error) {
	var m model.Message
	if err := ms.db.Model(&model.Message{}).
		Where("id = ?", id).First(&m).Error; err != nil {
		return &m, err
	}

	return &m, nil
}

func (ms *messageService) List(messageType string, page model.Page) (int64, []model.Message, error) {
	var count int64
	var messages []model.Message

	if messageType == "" {
		if err := ms.db.Model(&model.Message{}).
			Count(&count).
			Offset(page.Offset).
			Limit(page.PageSize).
			Order("update_time desc").
			Find(&messages).Error; err != nil {
			return count, messages, err
		}
	} else {
		if err := ms.db.Model(&model.Message{}).
			Where("message_type = ?", messageType).
			Count(&count).
			Offset(page.Offset).
			Limit(page.PageSize).
			Order("update_time desc").
			Find(&messages).Error; err != nil {
			return count, messages, err
		}
	}

	return count, messages, nil
}
