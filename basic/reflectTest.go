package basic

import (
	"fmt"
	"reflect"
)

type myInt int64

type person struct {
	name string
	age  int
}
type book struct {
	title string
}

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type Interfaces interface {
	inter() int
}

func (s student) Study() string {
	msg := "good good study,day day up!"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "good good sleep,quick quick grep up"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println("This interface have ", t.NumMethod(), "method")
	//:This interface have  2 method
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		//:method name:Sleep
		//:method name:Study
		fmt.Printf("method:%s\n", methodType)
		//:method:func() string
		//:method:func() string
		var args = []reflect.Value{}
		v.Method(i).Call(args)
		//:good good sleep,quick quick grep up
		//:good good study,day day up!
	}
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func function() bool {
	return true
}

func ReflectTest() {
	var a float32 = 3.14
	reflectType(a)
	//:type:float32 kind:float32
	reflectValue(a)
	//:type is float32, value is 3.140000
	var b int64 = 100
	reflectType(b)
	//:type:int64 kind:int64
	reflectValue(b)
	//:type is int64, value is 100
	re := reflect.ValueOf(10)
	fmt.Printf("type re : %T\n", re) // type re :reflect.Value
	//:type re :reflect.Value
	reflectType(re)
	//:type:Value kind:struct

	//2.array,slice,map,pointer... reflect.Name return nil
	var c *float32
	var d myInt
	var e rune
	reflectType(c)
	//:type: kind:ptr
	reflectType(d)
	//:type:myInt kind:int64
	reflectType(e)
	//:type:int32 kind:int32

	var f = person{
		name: "wardrobe",
		age:  18,
	}
	var g = book{title: "BaGa"}
	reflectType(f)
	//:type:person kind:struct
	reflectType(g)
	//:type:book kind:struct

	var s = []int{1, 2, 3, 4, 5}
	reflectType(s)
	//:type: kind:slice
	var m = map[string]int{"goku": 9001, "gohan": 2044}
	reflectType(m)
	//:type: kind:map
	var arr = [3]int{1, 2, 3}
	reflectType(arr)
	//:type: kind:array

	//3.Modify the value of a variable through reflection
	var q int64 = 100
	reflectSetValue(&q)
	reflectValue(q)
	//:type is int64, value is 200

	//4.reflect.ValueOf(Variables).IsNil()
	//Variables must be the following: channels, functions, interfaces, maps, pointers, slices
	var r *int
	fmt.Println("var r *int IsNil:", reflect.ValueOf(r).IsNil())
	//:var r *int IsNil: true
	var r1 chan int
	fmt.Println("var r1 chan IsNil:", reflect.ValueOf(r1).IsNil())
	//:var r1 chan IsNil: true
	fmt.Println("var function IsNil:", reflect.ValueOf(function).IsNil())
	//:var function IsNil: false
	//var r2 Interfaces
	//fmt.Println("var interface IsNil:", reflect.ValueOf(r2).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	//:nil IsValid: false
	z := struct{}{}
	fmt.Println("not exist struct field:", reflect.ValueOf(z).FieldByName("abc").IsValid())
	//:not exist struct field: false
	fmt.Println("not exit struct method:", reflect.ValueOf(z).MethodByName("abc").IsValid())
	//:not exit struct method: false
	// map
	t := map[string]int{}
	fmt.Println("map not exist this key:", reflect.ValueOf(t).MapIndex(reflect.ValueOf("fat")).IsValid())
	//:map not exist this key: false

	//5.through reflect get struct value
	stu1 := student{
		Name:  "john",
		Score: 90,
	}

	stu := reflect.TypeOf(stu1)
	fmt.Println(stu.Name(), stu.Kind())
	//:student struct
	for i := 0; i < stu.NumField(); i++ {
		field := stu.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
		//name:Name index:[0] type:string json tag:name
		//name:Score index:[1] type:int json tag:score
	}

	if scoreField, ok := stu.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
		//name:Score index:[1] type:int json tag:score
	}

	printMethod(stu1)
}
