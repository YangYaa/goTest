package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var (
	//Gauge 仪表盘类型
	opsQueued = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "SRG:Total",
		Help: "TestGauge HELP",
	})
	//Gauge 仪表盘类型
	opsQueued2 = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "SRG:TotalLink",
		Help: "TestGauge HELP",
	})
)

func PrometheusClient() {

	// 创建自定义注册表
	registry := prometheus.NewRegistry()

	registry.MustRegister(opsQueued)
	registry.MustRegister(opsQueued2)

	go func() {
		for {
			opsQueued.Add(15)
			time.Sleep(10 * time.Second)
		}
	}()

	go func() {
		for {
			opsQueued2.Add(15)
			time.Sleep(10 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}))
	//http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8088", nil))
}
