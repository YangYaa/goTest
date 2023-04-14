package work

import (
	"fmt"
	"net"
)

type Message struct {
	MsgNum   int    `json:"msgNum"`
	TotalNum int    `json:"totalNum"`
	Message  string `json:"message"`
}

func SimSendIo(sendcontent []byte, ipaddr string) {

	conn, err := net.Dial("tcp", ipaddr)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	conn.Write(sendcontent)
}

func SimReceiveIo() (recContent []byte) {
	conn, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer conn.Close()

	for {
		conn, err := conn.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			continue
		}
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		recContent = buf[:n]
		return
	}
}
