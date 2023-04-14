package goWebSocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type wsConn = websocket.Conn

func InitWebSocket() {
	//int a 20 size receive upload message chan
	init_mes_chan()
	LogWsMgr = wsMgr{
		Users:      make(map[*wsFd]bool),
		Broadcase:  make(chan []byte),
		Register:   make(chan *wsFd),
		Unregister: make(chan *wsFd),
	}
	//listen channel message and send to web client
	go LogWsMgr.Start(RecvMesChan)
}

func UpgradeFromGin(c *gin.Context) (*wsConn, error) {
	return (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}).Upgrade(c.Writer, c.Request, nil)
}

func RunWsGinServer(wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()
	var port uint16
	port = 9990
	return NewGinServer("/kong/test").AddService("wstest", wsTest).AddService("logup", southRoute).Run("192.168.50.211", port)
}
