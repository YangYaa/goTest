package basic

import (
	"fmt"
)

// interface definition
type VowelsFinder interface {
	//rune = int32
	FindVowels() []rune
}

type MyString string

// MyString implements VowelsFinder
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
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)
	//:Type = struct { name string }, value = {Naveen R}

	//3.type assert
	var t interface{} = 56
	assert(t)
	var u interface{} = "Steven Paul"
	assert(u)
}

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}
