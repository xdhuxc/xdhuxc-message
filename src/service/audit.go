package service

import (
	"github.com/jinzhu/gorm"
	"github.com/xdhuxc/xdhuxc-message/src/model"
	"time"
)

type auditService struct {
	db *gorm.DB
}

func newAuditService(db *gorm.DB) *auditService {
	return &auditService{db}
}

func (as *auditService) Create(oa model.OperationAudit) error {
	oa.CreateTime = time.Now()
	return as.db.Create(&oa).Error
}

func (as *auditService) List(page model.Page) (int64, []model.OperationAudit, error) {
	var oas []model.OperationAudit
	var count int64

	if err := as.db.Model(&model.OperationAudit{}).
		Count(&count).
		Order("create_time desc").
		Order("id desc").
		Limit(page.PageSize).
		Offset(page.Offset).Find(&oas).Error; err != nil {
		return count, oas, err
	}

	return count, oas, nil
}

func (as *auditService) FindByUser(user string, page model.Page) (int64, []model.OperationAudit, error) {
	var oas []model.OperationAudit
	var count int64

	if err := as.db.Model(&model.OperationAudit{}).
		Where("user = ?", user).
		Count(&count).
		Order("create_time desc").
		Limit(page.PageSize).
		Offset(page.Offset).Find(&oas).Error; err != nil {
		return count, oas, err
	}

	return count, oas, nil
}
