package rpc

import (
	"fmt"
	"log"
	"net/rpc"
)

func RpcClient() {
	// build tcp connection
	client, err := rpc.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error:", err)
	}
	fmt.Printf("ServiceA.Add: %d+%d=%d\n", args.X, args.Y, reply)

	// Asynchronous call
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replyCall := <-divCall.Done
	fmt.Println("IS Have Error:", replyCall.Error)
	fmt.Println("Result       :", reply2)
}
