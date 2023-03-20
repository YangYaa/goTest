package basic

import (
	"fmt"
)

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

func Accumulate(value int) func() int {
	// return a closure
	return func() int {
		// accumulation
		value++
		// return a accumulation value
		return value
	}
}
func GoClosure() {

	RegistRandomAccessExecCb(randomAccess)
	// create an accumulator
	accumulator := Accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Printf("%p\n", &accumulator)

	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", &accumulator2)
}
