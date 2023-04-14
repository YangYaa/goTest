package basic

import (
	"flag"
	"fmt"
)

var addr_client_s = flag.String("d", "192.168.50.211:8080", "http service address")

func GoFlagTest() {

	flag.Parse()
	fmt.Println("addr_client_s:", *addr_client_s)
}
