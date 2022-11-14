package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TestJson struct {
	Name string `json:name`
}

var Json = new(TestJson)

func LoadJsonFile() {
	data, err := ioutil.ReadFile("../json/test.json")
	if err != nil {
		fmt.Println("Read file error: ", err)
		return
	}
	json.Unmarshal(data, Json)
	fmt.Println("The read json is ", Json)
}
