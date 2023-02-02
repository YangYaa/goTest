package basic

import "fmt"

func SliceTest() {
	s := []int{1, 3, 5}

	//切片的遍历
	for index, value := range s {
		fmt.Println(index, value)
	}
}
