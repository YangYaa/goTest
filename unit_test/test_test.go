package unit_test

import (
	"fmt"
	"goTest/3rd/gorilla"
	"goTest/basic"
	"goTest/channel"
	"goTest/gin"
	"goTest/goHeap"
	"goTest/goIcmp"
	"goTest/goWebSocket"
	"goTest/goroutine"
	"goTest/http"
	"goTest/http/router"
	"goTest/json"
	"goTest/panicRecover"
	"goTest/prometheus"
	"goTest/rpc"
	"goTest/sync"
	"goTest/sync/errorGroup"
	"goTest/sync/mapManage"
	"goTest/uploadFile"
	"goTest/work"
	"testing"
	"time"
)

// go test -v -test.run

func TestIcmp(t *testing.T) {
	goIcmp.IcmpTest()
}

func TestGoFlag(t *testing.T) {
	basic.GoFlagTest()
}

func TestGorillaServer(t *testing.T) {
	gorilla.GorillaServer()
}

func TestGorillaClient(t *testing.T) {
	gorilla.GorillaClient()
}

func TestWebWS(t *testing.T) {
	goWebSocket.GoWebSocket2()
	//goWebSocket.NewGinServer("/kong/test")
}

func TestNewGinServer(t *testing.T) {
	goWebSocket.GinMain()
	//goWebSocket.NewGinServer("/kong/test")
}

func TestGoHeap2(t *testing.T) {
	goHeap.GoHeapTest2()
}

func TestGoHeap(t *testing.T) {
	goHeap.GoHeapTest()
}

func TestChan(t *testing.T) {
	basic.UnbuffChannel()
}

func TestGoRoutine1(t *testing.T) {
	basic.GoRoutineTest()
	basic.GoRoutineTest2()
}

func TestCallBackFunction2(t *testing.T) {
	basic.CallBackFunction2Test()
}

func TestCallBackFunction(t *testing.T) {
	basic.CallBackFunctionTest()
}

func TestTcpDecodeClient(t *testing.T) {
	basic.TcpClientEncodeTest()
}

func TestTcpEncodeServer(t *testing.T) {
	basic.TcpServerDecodeTest()
}

func TestTcpClient(t *testing.T) {
	basic.TcpClientTest()
}

func TestTcpServer(t *testing.T) {
	basic.TcpServerTest()
}

func TestGoReflect(t *testing.T) {
	basic.ReflectTest()
}
func TestGoArray(t *testing.T) {
	basic.ArrayTest()
}

func TestGoInterFace(t *testing.T) {
	basic.InterFaceTest()
}

func TestGoPointer(t *testing.T) {
	basic.PointerTest()
}

func TestGoCateLog(t *testing.T) {
	basic.CateLog()
}
func TestGoMap(t *testing.T) {
	basic.MapMake()
}

func TestGoClosure(t *testing.T) {
	basic.GoClosure()
}

func TestSimSend(t *testing.T) {
	work.SimSendMsg()
}

func TestSimRecv(t *testing.T) {
	work.SimRecvMsg()
}

func TestGoLog(t *testing.T) {
	basic.Log()
}
func TestRpcService(t *testing.T) {
	rpc.Service()
}

func TestRpcClient(t *testing.T) {
	rpc.Client()
}

func TestGoWebSocket(t *testing.T) {
	goWebSocket.GoWebSocket()
}

func TestPanicTest(t *testing.T) {
	panicRecover.PanicTest()
}

func TestWaitChan(t *testing.T) {
	channel.WaitChan()
}

func TestGinUpload(t *testing.T) {
	uploadFile.GinUpload()
}

func TestGinUploadClient(t *testing.T) {
	uploadFile.GinUploadClient()
}

func TestWaitUpLoadFile(t *testing.T) {
	uploadFile.UpLoadFile()
}

func TestHttpServer(t *testing.T) {
	//gin.InitialModel()
	http.HttpServer()
}

func TestMapMake(t *testing.T) {
	basic.MapMake()
}

func TestSliceTest(t *testing.T) {
	basic.SliceTest()
}

func TestGoRoutineTest(t *testing.T) {
	goroutine.GoRoutineTest()
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

func TestSyncOnceTest(t *testing.T) {
	sync.SyncOnceTest()
}

func TestSyncMapSafeTest(t *testing.T) {
	sync.SyncMapSafeTest()
}

func SyncMapUnSafeTest(t *testing.T) {
	sync.SyncMapUnSafeTest()
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
