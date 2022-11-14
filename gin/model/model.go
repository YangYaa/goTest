package model

type DbModel interface {
	Query() (profileList []DbModel, err error)
}
