package basic

import (
	"fmt"
)

// VowelsFinder interface definition
type VowelsFinder interface {
	// FindVowels rune = int32
	FindVowels() []rune
}

type WashingMachine interface {
	wash()
	dry()
}

// Payer interface definition
type Payer interface {
	Pay(int64)
}

type Interface interface {
	export()
}

type ZhiFuBao struct {
}

type WeChat struct {
}

type dryer struct{}

type haier struct {
	dryer
	//embed dryer
}

type MyString string

// M1 interface
type M1 interface {
	Say()
}

// M2 interface
type M2 interface {
	Move()
}

type Dog struct {
	Name string
}

type Car struct {
	Brand string
}

type reverse struct {
	Interface
}

type IntSlice []int

// Export attaches the methods of Interface
type Export int

func (x Export) export() { fmt.Println("The export merchandise is jar") }

func (r reverse) export() {
	fmt.Println("The export merchandise is tobacco")
}

func (x IntSlice) Len() int           { return len(x) }
func (x IntSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func Reverse(data Interface) Interface {
	return &reverse{data}
}

func (d dryer) dry() {
	fmt.Println("dry function")
}

func (h haier) wash() {
	fmt.Println("wash function")
}

func (d Dog) Say() {
	fmt.Printf("%s wang!wang!wang!\n", d.Name)
}

func (d Dog) Move() {
	fmt.Printf("%s can move\n", d.Name)
}

func (c Car) Move() {
	fmt.Printf("%s speed is 70 miles\n", c.Brand)
}

// Pay WeChat Method
func (w *WeChat) Pay(amount int64) {
	fmt.Printf("use wechat pay：%.2f RMB\n", float64(amount))
}

// Pay ZhiFubao Method
func (z *ZhiFuBao) Pay(amount int64) {
	fmt.Printf("use zhifubao pay %.2f RMB\n", float64(amount))
}

// Checkout Method
func Checkout(obj Payer) {
	// 支付100元
	obj.Pay(100)
}

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

	//4.Test Pay Method
	Checkout(&ZhiFuBao{})
	//:use zhifubao pay 100.00 RMB
	Checkout(&WeChat{})
	//:use wechat pay：100.00 RMB

	//5.a interface variable can storage all implement this interface variable
	var x Payer
	y := ZhiFuBao{}
	z := WeChat{}
	x = &y
	x.Pay(100)
	//:use zhifubao pay 100.00 RMB
	x = &z
	x.Pay(100)
	//:use wechat pay：100.00 RMB

	//6.one type can implement multiple interface
	var d = Dog{Name: "doudou"}
	var say M1 = d
	var move M2 = d
	say.Say()
	//:doudou wang!wang!wang!
	move.Move()
	//:doudou can move

	//7.one interface can implement by multiple type
	var obj = Car{Brand: "BMW"}
	obj.Move()
	//:BMW speed is 70 miles

	//8.All methods in an interface do not necessarily need to be implemented by one type
	//This can be achieved by embedding other types or structures in the type
	var embed WashingMachine
	haier := haier{}
	embed = haier
	embed.dry()
	//:dry function
	embed.wash()
	//:wash function

	//9.Embedding interfaces in structures
	/*----------test code----------
	initSlice := IntSlice{3, 1, 5, 7, 8}
	fmt.Println(sort.Reverse(initSlice).Less(3, 1))
	sort.Reverse(initSlice).Swap(3, 1)
	fmt.Println(initSlice)
	----------test code----------*/
	var ex Export
	ex.export()
	//:The export merchandise is jar
	Reverse(ex).export()
	//:The export merchandise is tobacco

	//10.Empty interface
	var emptyInterface interface{}
	emptyInterface = "hello"
	fmt.Printf("type:%T value:%v\n", emptyInterface, emptyInterface)
	//:type:string value:hello
	emptyInterface = 100
	fmt.Printf("type:%T value:%v\n", emptyInterface, emptyInterface)
	//:type:int value:100
	emptyInterface = true
	fmt.Printf("type:%T value:%v\n", emptyInterface, emptyInterface)
	//:type:bool value:true
	emptyInterface = Dog{}
	fmt.Printf("type:%T value:%v\n", emptyInterface, emptyInterface)
	//:type:basic.Dog value:{}
	show("hello")
	//:type:string value:hello
	//empty interface as map value
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "john"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
	//:map[age:18 married:false name:john]
}

// empty interface as function parameters
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

// Type Asserts
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
