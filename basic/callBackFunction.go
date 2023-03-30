package basic

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type callback func(data []byte, err error)

var wg sync.WaitGroup

func fetch(url string, c callback) {
	go func() {
		// 发送HTTP GET请求
		resp, err := http.Get(url)
		if err != nil {
			c(nil, err)
			return
		}
		defer resp.Body.Close()

		// 读取响应数据
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c(nil, err)
			return
		}
		time.Sleep(10 * time.Second)
		// 调用回调函数，传递响应数据和错误信息
		c(data, nil)
		wg.Done()
	}()
}

var RandomAccessExecCallBack func(gnbUuId int) int = nil

func CallRaExecCb(cb func(int) int, gnbUuId int) int {
	return cb(gnbUuId)
}

func randomAccess(gnbUuId int) int {
	fmt.Println("The uuid is ", gnbUuId)
	return 0
}

func RaExecRegist(cb func(int) int) {
	RandomAccessExecCallBack = cb
}

func RegistRandomAccessExecCb(cb func(int) int) {
	RaExecRegist(func(gnbUuId int) int {
		fmt.Println("The gnbUuId is", gnbUuId)
		return CallRaExecCb(cb, gnbUuId)
	})
}

func CallBackFunctionTest() {
	url := "https://www.baidu.com"
	wg.Add(1)
	fetch(url, func(data []byte, err error) {
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(string(data))
	})
	fmt.Println("Waiting for response...")
	wg.Wait()
}

func CallBackFunction2Test() {
	RegistRandomAccessExecCb(randomAccess)
	RandomAccessExecCallBack(10)
}
