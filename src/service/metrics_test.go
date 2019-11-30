package service

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var ms *metricsService

func init() {
	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"wh19940423",
		"127.0.0.1:3306",
		"xdhuxc-message")
	gdb, err := gorm.Open("mysql", uri)
	if err != nil {
		log.Fatalln(err)
	}
	err = gdb.DB().Ping()
	if err != nil {
		log.Errorln(err)
	}

	ms = newMetricsService(gdb)
}

func TestMetricsService_Count(t *testing.T) {
	mr := ms.Count()
	t.Log(mr.String())
}
