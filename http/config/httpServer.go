package config

import (
	"errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

type HttpServer struct {
	server http.Server
	g      *errgroup.Group
}

//initial http.Server
func NewHttpServer() *HttpServer {
	return new(HttpServer)
}

func (s *HttpServer) Start() {
	s.g.Go(func() error {
		return s.server.ListenAndServe()
	})
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

func (s *HttpServer) HttpHandler(r http.Handler) {
	s.server.Handler = r
}
