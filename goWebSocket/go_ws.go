package goWebSocket

import (
	"html/template"
	"log"
	"net/http"
)

func WebSocketWebTest(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../goWebSocket/wsweb.html")
	t.Execute(rw, nil)
}

func GoWebSocket2() {
	http.HandleFunc("/", WebSocketWebTest)
	if err := http.ListenAndServe("192.168.50.211:9992", nil); err != nil {
		log.Fatal(err)
	}
}
