package http

import (
	"fmt"
	conf "goTest/http/config"
	"goTest/http/router"
)

func HttpServer() {
	//initial http server config
	config := conf.GetConfigInstance()
	httpserver := conf.NewHttpConfig(config)
	//initial http server
	server := conf.NewHttpServer()
	//initial http handler
	server.HttpHandler(router.NewRouter(router.NewRoutes()))
	//because HttpServer implement Initialize and Start function,so it can add to ServerFunc
	httpserver.AddServer(server)
	err := httpserver.Initialize()
	if err != nil {
		fmt.Println("failed to initialize error")
	}
	httpserver.Start()
}
