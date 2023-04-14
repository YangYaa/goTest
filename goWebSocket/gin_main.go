package goWebSocket

import (
	"sync"
)

func GinMain() {
	wg := &sync.WaitGroup{}
	//init web socket channel
	InitWebSocket()
	RunWsGinServer(wg)
	wg.Wait()
}
