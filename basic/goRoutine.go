package basic

import "fmt"

func goodbye() {
	fmt.Println("goodbye son")
	wg.Done()
	//Tell the current goroutine to complete
}

func goodbye2(i int) {
	defer wg.Done()
	fmt.Println("goodbye son", i)
}

func recvChan(c chan int) {
	fmt.Println("this goroutine will block,still main process send data to chan")
	ret := <-c
	fmt.Println("recv chan success!", ret)
}

// 1.If you don't care about the results of concurrent operations or
// if there are other ways to collect the results of concurrent operations
func GoRoutineTest() {
	wg.Add(1)
	//Register a goroutine
	go goodbye()
	fmt.Println("goodbye dad")
	wg.Wait()
	//Blocking goroutine waiting for registration to complete
}

// 2.Start multiple goroutines
func GoRoutineTest2() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // Start a goroutine to register +1
		go goodbye2(i)
	}
	wg.Wait()
	//Waiting for all registered goroutines to end

	for i := 0; i < 5; i++ {
		fmt.Println("why?", i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}

// 3.Unbuffered channel
func UnbuffChannel() {
	ch := make(chan int)
	go recvChan(ch)
	//Create a goroutine to receive values from a channel
	ch <- 10
	fmt.Println("send success!")
}
