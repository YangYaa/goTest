package work

import (
	"encoding/json"
	"fmt"
)

func SimSendMsg() {
	send := new(Message)
	send.MsgNum = 1
	send.TotalNum = 10
	send.Message = "test"
	content, err := json.Marshal(send)
	if err != nil {
		fmt.Println("The marshal json error")
		return
	}
	//用户将定义好的content消息提供给send接口发出
	SimSendIo(content, "127.0.0.1:9091")
}
