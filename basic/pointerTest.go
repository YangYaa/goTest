package basic

import "fmt"

func PointerTest() {
	//1.init pointer
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	//:type of "a" is *int
	fmt.Println("address of b is", a)
	//:address of b is 0x1040a124

	//2.through a new() init a pointer
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	//:Size value is 0, type is *int, address is 0x414020

	//3.function return local variable pointer is valid
	d := hello()
	fmt.Println("Value of d", *d)
	//:Value of d 5
}

func hello() *int {
	i := 5
	return &i
}
