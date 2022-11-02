package unit_test

import (
	"goTest/channel"
	"goTest/http"
	"goTest/http/router"
	"goTest/sync/errorGroup"
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

func TestErrorGroupNotBreak(t *testing.T) {
	errorGroup.ErrorGroupNotBreak()
}
func TestErrorGroupBreak(t *testing.T) {
	errorGroup.ErrorGroupBreak()
}
