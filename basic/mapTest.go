package basic

import "fmt"

func MapMake() {
	
	//1.init a map
	//init through map keyword
	lookup := map[string]int{"goku": 9001, "gohan": 2044}
	//init through make function
	//init a new map must apply for memory through make
	cMap := make(map[string]int)
	cMap["Beijing"] = 1
	fmt.Println("lookup:", lookup)
	//:lookup: map[gohan:2044 goku:9001]
	fmt.Println("cMap:", cMap)
	//:cMap: map[Beijing:1]

	//2.search value among map
	cities := map[string]int{
		"Beijing": 100000,
		"Hunan":   430000,
	}
	city := "Beijing"
	postCode := cities[city]
	fmt.Println("City:", city, "PostCode:", postCode)
	//:City:Beijing PostCode: 100000
	fmt.Println("City:", cities["Shanghai"])
	//:City:0

	//3.search key among map
	newEmp := "Guangzhou"
	value, ok := cities[newEmp]
	if ok == true {
		fmt.Println("PostCode：", value)
		return
	}
	fmt.Println(newEmp, "PostCode is not Exit")
	//:Guangzhou PostCode is not Exit

	//4.range map
	for key, value := range cities {
		fmt.Printf("cities[%s] = %d\n", key, value)
	}
	//:cities[Beijing] = 100000
	//:cities[Hunan] = 430000

	//5.delete map key
	fmt.Println("map before deletion", cities)
	//:map before deletion map[Beijing:100000 Hunan:430000]
	delete(cities, "Beijing")
	//go delete function not have return value,even we delete a not exit key,it will not report error
	fmt.Println("map after deletion", cities)
	//:map after deletion map[Hunan:430000]
}
