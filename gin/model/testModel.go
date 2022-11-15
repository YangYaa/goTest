package model

import (
	"goTest/mysql"
)

type TestDb struct {
	Id   *uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name *string `json:"name" gorm:"column:name;uniqueIndex;type:varchar(32);not null"`
}

func (p *TestDb) Query() (profileList []DbModel, err error) {
	var list []TestDb
	err = mysql.GetMysqlInstance().Where(p).Find(&list).Error
	for _, val := range list {
		t := val
		profileList = append(profileList, &t)
	}
	return profileList, err
}
