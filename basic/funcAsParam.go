package basic

import "fmt"

//定义函数类型 op
type op func(a, b int) int

func add(a, b int) int {
	return a + b
}

//sub作为函数名可以看成是 op 类型的常量
func sub(a, b int) int {
	return a - b
}

//形参指定传入参数为函数类型op
func Oper(fu op, a, b int) int {
	return fu(a, b)
}
func funcAsParam() {
	//在go语言中函数名可以看做是函数类型的常量，所以我们可以直接将函数名作为参数传入的函数中
	a := Oper(add, 1, 2)
	fmt.Println(a)
	b := Oper(sub, 1, 2)
	fmt.Println(b)
}
