package basic

import "fmt"

func MapMake() {
	//new(T) 函数是一个分配内存的内置函数，为每个类型分配一片内存，并初始化为零值且返回其内存地址
	var v *int
	//& 是取地址符号 , 即取得某个变量的地址 v的地址: 0xc00040a200 v指向的地址: <nil>
	fmt.Println("v的地址:", &v, "v指向的地址:", v)
	v = new(int)
	//v的地址: 0xc000010040 v指向的地址: 0xc00040a200 v指向的地址的值: 0
	fmt.Println("v的地址:", &v, "v指向的地址:", v, "v指向的地址的值:", *v)
	//make函数是Go的内置函数，它的作用是为slice、map或chan初始化并返回引用
	scoreMap := make(map[string]int)
	//声明的时候不需要知道 map 的长度，map 是可以动态增长的,未初始化的 map 的值是 nil
	fmt.Printf("scoreMap的地址:%T\n", scoreMap)
	scoreMap["张三"] = 90
	fmt.Println("len scoreMap", len(scoreMap))
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}
