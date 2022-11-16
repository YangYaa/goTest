package config

import (
	"golang.org/x/sync/errgroup"
	"log"
)

var (
	g errgroup.Group
)

type ServerFunc interface {
	Initialize(data ...interface{}) error
	Start()
}

type Http struct {
	Servers []ServerFunc  //接口数组
	Config  *GlobalConfig //http配置
}

//initial http config
func NewHttpConfig(config *GlobalConfig) *Http {
	return &Http{Config: config}
}

func (app *Http) Start() {
	var err error
	for _, s := range app.Servers {
		s.Start()
	}
	if err = g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func (app *Http) Initialize() error {
	for index, s := range app.Servers {
		err := s.Initialize(&g, app.Config, index)
		if err != nil {
			return err
		}
	}
	return nil
}

func (app *Http) AddServer(s ServerFunc) {
	app.Servers = append(app.Servers, s)
}
