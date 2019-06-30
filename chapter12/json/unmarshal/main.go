package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age int
	Skill string
}

//演示将json字符串，反序列化成struct
func unmarshalStruct() {
	jsonStr := "{\"Name\":\"小明\",\"Age\":20,\"Skill\":\"python\"}"
	var person Person
	err := json.Unmarshal([]byte(jsonStr),&person)
	if err != nil {
		fmt.Println("unmarshal err=",err)
	}
	fmt.Println("反序列化后person=",person)
}

//演示将json字符串反序列化成map
func unmarshalMap()  {
	jsonStr := "{\"Name\":\"小明\",\"Age\":20,\"Skill\":\"python\"}"
	var m map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr),&m)
	if err != nil {
		fmt.Println("unmarshal err=",err)
	}
	fmt.Println("反序列化后m=",m)
}

//演示将json字符串反序列化成切片
func unmarshalSlice() {
	jsonStr := "[{\"Name\":\"小明\",\"Age\":20,\"Skill\":\"python\"},{\"Name\":\"小华\",\"Age\":21,\"Skill\":\"java\"}]"
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr),&slice)
	if err != nil {
		fmt.Println("unmarshal err=",err)
	}
	fmt.Println("反序列化后slice=",slice)


}


func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}