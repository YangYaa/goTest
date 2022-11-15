package json

import (
	"encoding/json"
	"fmt"
	"goTest/3rd/pretty"
	"io/ioutil"
)

type TestJson struct {
	Name  string `json:name`
	Array []int  `json:array`
}

var Json = new(TestJson)

func LoadJsonFile() {
	//读取文件中的数据
	data, err := ioutil.ReadFile("../json/test.json")
	if err != nil {
		fmt.Println("Read file error: ", err)
		return
	}
	json.Unmarshal(data, Json)
	fmt.Println("The read json is ", Json)
	//追加数据到结构体数组中
	Json.Array = append(Json.Array, 99)
	dataJson, _ := json.Marshal(Json)
	ioutil.WriteFile("../json/test.json", pretty.Pretty(dataJson), 0777)
}
