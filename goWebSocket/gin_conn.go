package goWebSocket

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

type Server struct {
	*gin.Engine
	prefix string
}
type UriType = string
type methodType = uint8
type Route map[UriType]map[methodType]gin.HandlerFunc

const (
	GET uint8 = iota
	HEAD
	POST
	PUT
	PATCH
	DELETE
	CONNECT
	OPTIONS
	TRACE
)

func NewGinServer(prefix string) *Server {
	var engine *gin.Engine
	if "1" == os.Getenv("GIN_DEFAULT_ENGINE") {
		engine = gin.Default()
	} else {
		engine = gin.New()
		engine.Use(gin.Recovery())
	}
	engine.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "X-Requested-With"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           86400,
	}))
	return &Server{engine, prefix}
}

func (s *Server) AddService(uri string, routes Route) *Server {
	g := s.Group(s.prefix).Group(uri)
	for uri2, group := range routes {
		for method, callback := range group {
			switch method {
			case GET:
				g.GET(uri2, callback)
			case POST:
				g.POST(uri2, callback)
			case HEAD, PUT, PATCH, DELETE, CONNECT, OPTIONS, TRACE:
				fmt.Println("NOT SUPPOST THIS METHOD")
			}
		}
	}
	return s
}

func (s *Server) Run(ip string, port uint16) error {
	return s.Engine.Run(fmt.Sprintf("%s:%d", ip, port))
}
