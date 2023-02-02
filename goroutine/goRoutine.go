package goroutine

import (
	"fmt"
	"sync"
)

//声明全局等待组变量
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("hello,admin", i)
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("SUCCESS 接收成功", ret)
}

func getChannel(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("SUCCESS RECEIVE CHANNEL v:%#v ok:%#v\n", v, ok)
	}
}

// Producer2 返回一个接收通道
//<- chan int // 只接收通道，只能接收不能发送
func Producer2() <-chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()

	return ch
}

// Consumer2 参数为接收通道
func Consumer2(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

//sync并发原语

func GoRoutineTest() {
	/*
		for i := 0; i < 10; i++ {
			wg.Add(1) // 启动一个goroutine就登记+1
			go hello(i)
		}

	*/

	//chan <- int // 只发送通道，只能发送不能接收
	//<- chan int // 只接收通道，只能接收不能发送
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	fmt.Println("ERROR channel is already close!")
	getChannel(ch)
	fmt.Println("SUCCESS main goroutine done!")
	ch2 := Producer2()
	res2 := Consumer2(ch2)
	fmt.Println(res2) // 25
	wg.Wait()         // 阻塞等待登记的goroutine完成
}
