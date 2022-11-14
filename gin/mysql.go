package gin

import (
	"fmt"
	"goTest/gin/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var SqlDb *gorm.DB

func InitialMysql() {
	fmt.Println("The Db gorm Test")
	err := Initialize("127.0.0.1", "root", "123456", "testdb")
	if err != nil {
		fmt.Println("Initialize mysql error")
	}
	// initial mysql table through gorm struct
	err = SqlDb.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.TestDb{})
	if err != nil {
		fmt.Println("The insert into database error")
	}
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
