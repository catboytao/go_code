package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Score float32
	Sex string
}

func (m Monster) Print() {
	fmt.Println("----start---")
	fmt.Println(m)
	fmt.Println("----end----")
}

func (m Monster) GetSum(n1,n2 int) int {
	return n1 + n2
}

func (m Monster) Set(name string,age int,score float32,sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

func TestStruct(a interface{}) {
	//获取reflect.Type
	typ := reflect.TypeOf(a)
	//获取reflect.Value
	val := reflect.ValueOf(a)
	
	kd := val.Kind()
	//如果传入的不是struct，就退出
	if kd != reflect.Struct{
		fmt.Println("expect struct")
		return
	}
	
	//获取该结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n",num)

	//遍历结构体的所有字段
	for i := 0;i < num; i++ {
		fmt.Printf("Field %d:值为%v\n",i,val.Field(i))
		//获取到struct标签，主要需要通过reflect.Type来获取
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d:tag为%v\n",i,tagVal)
		}
	}

	//获取到该结构体有多少个方法
	numMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n",numMethod)

	//var params []reflect.Value
	val.Method(1).Call(nil)

	var args []reflect.Value
	args = append(args,reflect.ValueOf(10))
	args = append(args,reflect.ValueOf(20))
	res := val.Method(0).Call(args)  // []reflect.Value
	fmt.Println("res=",res[0].Int())

}


func main() {
	monster := Monster{
		Name : "牛魔王",
		Age : 21,
		Score : 98.5,
		Sex : "male",
	}
	TestStruct(monster)

}