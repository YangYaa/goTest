package model

import (
	"fmt"
	"goTest/mysql"
)

type TestDb struct {
	Id   *uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name *string `json:"name" gorm:"column:name;uniqueIndex;type:varchar(32);not null"`
}

func (p *TestDb) Check() error {
	return nil
}

func (p *TestDb) Create() error {
	return mysql.GetMysqlInstance().Create(p).Error
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

func (p *TestDb) QuerySouth() (profileList []DbModel, err error) {
	var list []TestDb
	err = mysql.GetMysqlInstance().Where(p).Find(&list).Error
	for _, val := range list {
		t := val
		profileList = append(profileList, &t)
	}
	return profileList, err
}

func (p *TestDb) Update() error {
	var err error
	if p.Id != nil {
		result := mysql.GetMysqlInstance().Model(p).Where("id = ?", *p.Id).Updates(p)
		if result.RowsAffected > 0 {
			err = nil
		} else if result.Error != nil {
			err = result.Error
		} else {
			fmt.Errorf("failed to update id = %d\n", *p.Id)
		}
	} else {
		err = fmt.Errorf("faile to update TestDb")
	}
	return err
}

func (p *TestDb) Delete() error {
	return mysql.GetMysqlInstance().Delete(p).Error
}
