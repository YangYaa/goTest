package config

import (
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

type ServerFunc interface {
	Initialize(data ...interface{}) error
	Start()
}

type Http struct {
	Servers []ServerFunc
	Config  *GlobalConfig
}

type HttpServer struct {
	server http.Server
	g      *errgroup.Group
}

func NewHttp(config *GlobalConfig) *Http {
	return &Http{Config: config}
}

func NewHttpServer() *HttpServer {
	return new(HttpServer)
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
func (app *Http) Start() {
	var err error
	for _, s := range app.Servers {
		s.Start()
	}
	if err = g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func (s *HttpServer) Initialize(data ...interface{}) error {
	group, ok := data[0].(*errgroup.Group)
	if !ok {
		return errors.New("Initialize Httpserver Failed")
	}
	appConfig, ok := data[1].(*GlobalConfig)
	if !ok {
		return errors.New("Initialize appConfig Failed")
	}
	configIndex, ok := data[2].(int)
	if !ok {
		return errors.New("Initialize configIndex Failed")
	}
	s.g = group
	serverConfig := appConfig.Server[configIndex]
	s.server.Addr = serverConfig.IP + ":" + serverConfig.Port
	s.server.ReadTimeout = time.Duration(serverConfig.ReadTimeout) * time.Second
	s.server.WriteTimeout = time.Duration(serverConfig.WriteTimeout) * time.Second
	return nil
}
func (s *HttpServer) Start() {
	s.g.Go(func() error {
		return s.server.ListenAndServe()
	})

}

func (app *Http) AddServer(s ServerFunc) {
	app.Servers = append(app.Servers, s)
}

func (s *HttpServer) HttpHandler(r http.Handler) {
	s.server.Handler = r
}
