package unit_test

import (
	"goTest/basic"
	"goTest/channel"
	"goTest/gin"
	"goTest/http"
	"goTest/http/router"
	"goTest/json"
	"goTest/prometheus"
	"goTest/sync/errorGroup"
	"testing"
)

//  go test -v -test.run
func TestWaitChan(t *testing.T) {
	channel.WaitChan()
}

func TestHttpServer(t *testing.T) {
	gin.InitialModel()
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

func TestLoadJsonFile(t *testing.T) {
	json.LoadJsonFile()
}

func TestInitialModel(t *testing.T) {
	gin.InitialModel()
}

func TestPrometheusClient(t *testing.T) {
	prometheus.PrometheusClient()
}

func TestStructEmbedInterface(t *testing.T) {
	basic.StructEmbedInterface()
}
