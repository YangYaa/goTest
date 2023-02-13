package goWebSocket

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
	"time"
)

func welcome(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../goWebSocket/gowebsocket.html")
	t.Execute(rw, nil)
}

func GoWebSocket() {
	http.HandleFunc("/", welcome)
	http.Handle("/goWebSocketTest", websocket.Handler(Echo))
	if err := http.ListenAndServe("192.168.50.137:9998", nil); err != nil {
		log.Fatal(err)
	}
}

func Echo(w *websocket.Conn) {
	var error error
	var msg string
	msg = "100"
	for {

		/*var reply string
		fmt.Println("The pro is run before Receive message")

			if error = websocket.Message.Receive(w, &reply); error != nil {
				fmt.Println("不能够接受消息 error==", error)
				break
			}


		fmt.Println("The pro is run after Receive message")
		fmt.Println("能够接受到消息了--- ", reply)
		msg := "我已经收到消息 Received:" + reply
		//  连接的话 只能是   string；类型的啊
		fmt.Println("发给客户端的消息： " + msg)
		*/
		if error = websocket.Message.Send(w, msg); error != nil {
			fmt.Println("不能够发送消息 悲催哦")
			break
		}
		msg = msg + "-|"
		time.Sleep(5 * time.Second)
	}
}
