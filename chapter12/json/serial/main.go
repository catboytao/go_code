package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string
	Age int
	Birthday string
	Sal float64
	Skill string
}
//对结构体进行序列化
func testStruct() {
	monster := Monster{
		Name : "牛魔王",
		Age : 500,
		Birthday : "2011-11-11",
		Sal : 8000.0,
		Skill : "牛魔拳",
	}
	//将monster序列化
	data,err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误 err=",err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n",string(data))
}
//对map进行序列化
func testMap() {
	//定义一个map
	var m map[string]interface{} 
	//使用map,需要make
	m = make(map[string]interface{})
	m["name"] = "红孩儿"
	m["age"] = 20
	m["address"] = "洪崖洞"
	//将map进行序列化
	data,err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化错误 err=",err)
	}
	//输出map序列化的结果
	fmt.Printf("m序列化后=%v\n",string(data))
}
//对切片进行序列化,使用[]map[string]interface{}切片
func testSlice(){
	var slice []map[string]interface{}
	var m1,m2 map[string]interface{}
	m1 = make(map[string]interface{})
	m2 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 19
	m1["sex"] = "male"
	slice = append(slice,m1)
	m2["name"] = "marry"
	m2["age"] = 20
	m2["sex"] = "famale"
	slice = append(slice,m2)
	data,err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化错误 err=",err)
	}
	//输出map序列化的结果
	fmt.Printf("slice序列化后=%v\n",string(data))
}
//对基本数据类型进行序列化
func testFloat64(){
	var num1 float64 = 2345.56
	data,err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=",err)
	}
	//输出num1序列化的结果
	fmt.Printf("num1序列化后=%v\n",string(data))
}

func main() {
	//演示将结构体，map,切片进行序列化
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}