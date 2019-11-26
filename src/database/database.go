package database

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/xdhuxc/xdhuxc-message/src/model"
)

func NewMysqlClient(database model.Database) (*gorm.DB, error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		database.User,
		database.Password,
		database.Host,
		database.DatabaseName)
	gdb, err := gorm.Open("mysql", uri)
	if err != nil {
		return nil, err
	}

	err = gdb.DB().Ping()
	if err != nil {
		return nil, err
	}

	gdb.LogMode(database.Log)
	gdb.DB().SetMaxIdleConns(database.MaxIdleConns)
	gdb.DB().SetMaxOpenConns(database.MaxOpenConns)
	gdb.DB().SetConnMaxLifetime(time.Hour)

	return gdb, nil
}
