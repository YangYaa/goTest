package basic

import "fmt"

func SliceTest() {
	//method 1 to init new a slice
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
	//method 2 to init new a slice
	c := []int{6, 7, 8}
	fmt.Println(c)

	//2.operate on the slice ultimately affects the underlying array
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59} //array
	dsLice := darr[2:5]                                   //slice
	fmt.Println("array before", darr)
	//:array before [57 89 90 82 100 78 67 69 59]
	for i := range dsLice {
		dsLice[i]++
	}
	fmt.Println("array after", darr)
	//:array after  [57 89 91 83 101 78 67 69 59]

	//3.slice length: number of elements	map capacity: number of underlying array
	fruitArray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitSlice := fruitArray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitSlice), cap(fruitSlice))
	//:length of slice 2 capacity 6

	//4.reset slice content
	fruitSlice = fruitSlice[:cap(fruitSlice)]
	fmt.Println("After re-slicing length is", len(fruitSlice), "and capacity is", cap(fruitSlice))
	//:After re-slicing length is 6 and capacity is 6

	//5.Create a slice through make
	i := make([]int, 5, 5)
	fmt.Println(i)
	//:[0 0 0 0 0]

	//6.slice expansion principle
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Printf("The before address is %p\n", cars)
	//:0xc0005a45a0
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	//:cars: [Ferrari Honda Ford] has old length 3 and capacity 3
	cars = append(cars, "Toyota")
	//if the original slice is full,go will create a new array,length is twice the original,and copy the original element to the new
	//return this new slice quote
	fmt.Printf("The after address is %p\n", cars)
	//:0xc00060e060
	cars = append(cars, "RR")
	fmt.Printf("The after address is %p\n", cars)
	//:0xc00060e060
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))
	//:cars: [Ferrari Honda Ford Toyota RR] has new length 5 and capacity 6

	//7.slice as function value
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	//:slice before function call [8 7 6]
	subOne(nos)
	fmt.Println("slice after function call", nos)
	//:slice after function call [6 5 4]

	//8.Multidimensional slice
	pls := [][]string{
		{"C", "C++"},
		{"Java"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func subOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}

}
