package work

import (
	"encoding/json"
	"fmt"
)

func SimRecvMsg() {
	//调用recv io 获取 send io 发送的数据
	content := SimReceiveIo()
	message := new(Message)
	json.Unmarshal(content, message)

	fmt.Println("The message seq is ", message.MsgNum)
	fmt.Println("The message total num is ", message.TotalNum)
	fmt.Println("The message content is ", message.Message)
}
