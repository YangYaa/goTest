package basic

import "fmt"

func ArrayTest() {

	//1.init array
	var testArray [3]int
	// ... compiler will auto compute length
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"BJ", "SH", "SZ"}
	var indexArray = [...]int{1: 1, 3: 5}
	fmt.Println(testArray)
	//:[0 0 0]
	fmt.Println(numArray)
	//:[1 2 0]
	fmt.Println(cityArray)
	//:[BJ SH SZ]
	fmt.Println(indexArray)
	//:[0 1 0 5]

	//2.Two-dimensional array range
	a := [3][2]string{
		{"11", "12"},
		{"21", "22"},
		{"31", "32"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
	//:11      12
	//:21      22
	//:31      32

	//3.go array is not a quote type,is value type
	j := [3]int{10, 20, 30}
	modifyArray(j)
	//will modify a copy of j,not j itself
	fmt.Println(j) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b)
	//will modify a copy of b,not b itself
	fmt.Println(b) //[[1 1] [1 1] [1 1]]
}

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
