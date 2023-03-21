package rpc

import (
	"log"
	"net"
	"net/rpc"
)

func Service() {
	service := new(ServiceA)
	rpc.Register(service)
	l, e := net.Listen("tcp", ":9091")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		conn, _ := l.Accept()
		rpc.ServeConn(conn)
	}
}
