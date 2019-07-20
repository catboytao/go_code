package main

import (
	"fmt"
	"reflect"
)

//通过反射，修改

func reflect01(b interface{}) {
	//获取到reflect.Value
	rVal := reflect.ValueOf(b)
	//Elem返回rVal持有的接口保管的值的Value封装，或者rVal持有的指针指向的值的Value封装
	rVal.Elem().SetInt(20)
}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("num=",num)
}