package sync

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	once sync.Once
	wg   sync.WaitGroup
	m    = make(map[string]int)
	m2   = sync.Map{}
)

//确保某些操作即使在高并发的场景下也只会被执行一次

func SyncOnceTest() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// 这里要注意讲i显示的当参数传入内部的匿名函数
		go func(i int) {
			defer wg.Done()
			// fmt.Println("once", i)
			once.Do(func() {
				fmt.Println("once", i)
			})
		}(i)
	}

	wg.Wait()
	fmt.Printf("over")
}

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func SyncMapUnSafeTest() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func SyncMapSafeTest() {
	wg := sync.WaitGroup{}
	// 对m执行20个并发的读写操作
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)         // 存储key-value
			value, _ := m2.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
		//匿名函数调用方法相当于
		//f := func(n int){ xxxx }; f(i)
	}
	wg.Wait()
}
