package main

import (
	"fmt"
)

//定义一个Cat结构体
type Cat struct {
	Name string
	Age int
	Color string
}


func main() {
	//创建一个Cat对象
	var c Cat
	c.Name = "小白"
	c.Age = 3
	c.Color = "白色"
	fmt.Println("c=",c)
}