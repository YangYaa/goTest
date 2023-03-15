package basic

import (
	"fmt"
)

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
	// create an accumulator
	accumulator := Accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Printf("%p\n", &accumulator)

	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", &accumulator2)
}
