package basic

import (
	"errors"
	"fmt"
	"strings"
)

type calculation func(int, int) int

func Accumulate(value int) func() int {
	// return a closure
	return func() int {
		// accumulation
		value++
		// return a accumulation value
		return value
	}
}

func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

func addTest(x, y int) int {
	return x + y
}

func subTest(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return addTest, nil
	case "-":
		return subTest, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func addclosure() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func addclosure2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc2(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func GoClosure() {
	// create an accumulator
	accumulator := Accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Printf("%p\n", &accumulator)

	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", &accumulator2)

	//2.go function
	ret5 := intSum3(100)
	//:100 []
	ret6 := intSum3(100, 10)
	//:100 [10]
	ret7 := intSum3(100, 10, 20)
	//:100 [10 20]
	ret8 := intSum3(100, 10, 20, 30)
	//:100 [10 20 30]
	fmt.Println(ret5, ret6, ret7, ret8)
	//:100 110 130 160

	//3.define function type
	var c calculation
	c = addTest
	fmt.Printf("type of c:%T\n", c)
	//:type of c:basic.calculation
	fmt.Println(c(1, 2))
	//:3

	f := addTest
	fmt.Printf("type of f:%T\n", f)
	//:type of f:func(int, int) int
	fmt.Println(f(10, 20))
	//:30

	//4.functions as parameters
	ret2 := calc(10, 20, subTest)
	fmt.Println(ret2)
	//:-10

	//5.function as return value
	do("+")

	//6.anonymous function
	//save anonymous functions to variables
	adder := func(x, y int) {
		fmt.Println(x + y)
	}
	adder(10, 20)
	//Self-executing function: Anonymous function is directly executed after it is defined plus ()
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	//7.closure=Function+Reference Environment
	var fun = addclosure()
	fmt.Println(fun(10))
	//:10
	fmt.Println(fun(20))
	//:30
	fmt.Println(fun(30))
	//:60

	fun1 := addclosure()
	fmt.Println(fun1(40))
	//:40
	fmt.Println(fun1(50))
	//:90

	var fun2 = addclosure2(10)
	fmt.Println(fun2(10))
	//:20
	fmt.Println(fun2(20))
	//:40
	fmt.Println(fun2(30))
	//:70

	fun3 := addclosure2(20)
	fmt.Println(fun3(40))
	//:60
	fmt.Println(fun3(50))
	//:110

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt

	fun4, fun5 := calc2(10)
	fmt.Println(fun4(1), fun5(2))
	//:11 9
	fmt.Println(fun4(3), fun5(4))
	//:12 8
	fmt.Println(fun4(5), fun5(6))
	//:13 7
}
