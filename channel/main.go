package channel

import (
	"fmt"
	"time"
)

var ch chan int

func add_chan(){
	time.Sleep(10*time.Second)
	ch <- 1
}
func WaitChan() {
	ch = make(chan int, 1)
	go add_chan()

	for {
		select {
		case x := <-ch:
			fmt.Println(x)
			goto over
		default:
		}
	}
	over:
		fmt.Println("The pro is over")
}
