package model

type DbModel interface {
	Query() (profileList []DbModel, err error)
	Create() error
	Check() error
	Update() error
	Delete() error
}
