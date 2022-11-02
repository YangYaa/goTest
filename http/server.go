package http

import (
	"fmt"
	conf "goTest/http/config"
	"goTest/http/router"
)

func HttpServer() {
	config := conf.GetConfigInstance()
	httpserver := conf.NewHttp(config)
	server := conf.NewHttpServer()
	server.HttpHandler(router.NewRouter(router.NewRoutes()))
	httpserver.AddServer(server)
	err := httpserver.Initialize()
	if err != nil {
		fmt.Println("failed to initialize error")
	}
	httpserver.Start()
}
