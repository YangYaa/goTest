package basic

import (
	"fmt"
)

// VowelsFinder interface definition
type VowelsFinder interface {
	// FindVowels rune = int32
	FindVowels() []rune
}

type MyString string

// FindVowels MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func InterFaceTest() {

	//1.create interface and implement
	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())
	//:Vowels are [a e o]

	//2.use empty interface
	s := "Hello World"
	describe(s)
	//:Type = string, value = Hello World
	i := 55
	describe(i)
	//:Type = int, value = 55
	str := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(str)
	//:Type = struct { name string }, value = {Naveen R}

	//3.type assert
	var t interface{} = 56
	assert(t)
	//:x is a int is 56
	var u interface{} = "Steven Paul"
	assert(u)
	//:x is a string，value is Steven Paul
}

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assert(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
