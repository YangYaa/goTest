package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var SqlDb *gorm.DB

func GetMysqlInstance() *gorm.DB {
	return SqlDb
}

// initial mysql env
func Initialize(host, user, password, dbName string) error {
	var err error
	s := fmt.Sprintf("%s:%s@(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbName)
	SqlDb, err = gorm.Open(mysql.Open(s), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic("mysql connect error: = " + err.Error())
	}
	instance, _ := SqlDb.DB()
	instance.SetMaxIdleConns(10)
	instance.SetConnMaxLifetime(10 * time.Minute)
	instance.SetMaxOpenConns(100)
	return err
}
