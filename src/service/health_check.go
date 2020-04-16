package service

import (
	"github.com/jinzhu/gorm"
)

type healthCheckService struct {
	db *gorm.DB
}

func newHelathCheckService(db *gorm.DB) *healthCheckService {
	return &healthCheckService{db}
}

func (hcs *healthCheckService) Get() (model.HealthCheck, error) {
	return model.HealthCheck{
		Message: "It is ok!",
	}, nil
}
