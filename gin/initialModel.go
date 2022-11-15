package gin

import (
	"database/sql"
	"fmt"
	"goTest/gin/model"
	"goTest/mysql"
)

func InitialModel() {
	fmt.Println("The Db gorm Test")
	//create database if not exit
	dbTest, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	defer dbTest.Close()
	if err != nil {
		fmt.Println("Connect Mysql error")
	} else {
		createdb := "CREATE DATABASE testdb"
		dbTest.Exec(createdb)
	}

	//initial mysql connect
	err = mysql.Initialize("127.0.0.1", "root", "123456", "testdb")
	if err != nil {
		fmt.Println("Initialize mysql error")
	}
	//get mysql instance
	SqlDb := mysql.GetMysqlInstance()

	// initial mysql table through gorm struct
	err = SqlDb.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.TestDb{})
	if err != nil {
		fmt.Println("create table into database error", err)
	}
}
