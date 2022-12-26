package unit_test

import (
	"fmt"
	"goTest/basic"
	"goTest/channel"
	"goTest/gin"
	"goTest/http"
	"goTest/http/router"
	"goTest/json"
	"goTest/prometheus"
	"goTest/sync/errorGroup"
	"goTest/sync/mapManage"
	"testing"
	"time"
)

//  go test -v -test.run
func TestWaitChan(t *testing.T) {
	channel.WaitChan()
}

func TestHttpServer(t *testing.T) {
	gin.InitialModel()
	http.HttpServer()
}

func TestMapMake(t *testing.T) {
	basic.MapMake()
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

func TestAddInstance(t *testing.T) {
	mapInstance := mapManage.NewManager()
	mapInstance.AddInstance("test", "test.com")
	url := mapInstance.GetInstance("test")
	fmt.Println("The map get Instance url is ", url)
}

func TestPool(t *testing.T) {
	pools, err := channel.NewPool(2000)
	if err != nil {
		return
	}
	i := 0
	taskResultChan := make(chan *channel.TaskResult)
	defer close(taskResultChan)
	for i := 0; i < 10; i++ {
		task := channel.NewTask(channel.NotifyTask, []interface{}{"PUT", "callurl", i}, taskResultChan)
		err = pools.Put(task)
		if err == nil {
			i++
		}
	}
	complete := false
loop:
	for {
		select {
		case res := <-taskResultChan:
			{
				if res.OutPut != nil {
					v := res.OutPut.(string)
					if v == "notify successfully" {
						i--
						if i == 0 {
							complete = true
							break loop
						}
					}
				}
				if res.Err != nil {
					fmt.Println("have error message:", res.Err)
				}
			}
		case <-time.After(3 * time.Second):
			{
				fmt.Println("time out will break")
				break loop
			}
		}
	}
	if complete {
		fmt.Println("all the message is success")
	}

}
