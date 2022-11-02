package unit_test

import (
	"goTest/channel"
	"goTest/http"
	"goTest/http/router"
	"testing"
)

func TestWaitChan(t *testing.T) {
	channel.WaitChan()
}

func TestHttpServer(t *testing.T) {
	http.HttpServer()
}

func TestNewRouter(t *testing.T) {
	route := router.NewRoutes()
	router.NewRouter(route)
}
