package model

type TestDb struct {
	Id      *uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name    *string `json:"name" gorm:"column:name;uniqueIndex;type:varchar(32);not null"`
}

