package goWebSocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var wsTest = Route{
	"wsLog": {
		GET: WsLogTest,
	},
}

var southRoute = Route{
	"/ueLogUp": {
		POST: UeLogUpLoad,
	},
}

func WsLogTest(c *gin.Context) {
	//new a websocket connect
	conn, err := UpgradeFromGin(c)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		fmt.Println("Run websocket failed", err)
		return
	}
	fmt.Println("websocket client is start run :", c.ClientIP())
	LogWsMgr.RegClientUser(c.ClientIP(), conn)
}

func UeLogUpLoad(c *gin.Context) {
	if data, err := ioutil.ReadAll(c.Request.Body); err == nil {
		select {
		case RecvMesChan <- data:
			fmt.Println("log buff sent content is ", string(data))
		default:
			fmt.Println("log buffer is full we drop this log", string(data))
		}
	} else {
		fmt.Println("ue update log get data failed")
	}
}
