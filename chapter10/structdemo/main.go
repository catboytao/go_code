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
	//方式1
	var c Cat
	c.Name = "小白"
	c.Age = 3
	c.Color = "白色"
	fmt.Println("c=",c)
	//方式2
	var c2 Cat = Cat{"小花",5,"花色"}
	fmt.Println("c2=",c2)
	//方式3
	var c3 *Cat = new(Cat)
	(*c3).Name = "小黑"
	(*c3).Age = 2
	(*c3).Color = "黑色"
	fmt.Println("c3=",*c3)
	//方式4
	var c4 *Cat = &Cat{"小黄",3,"黄色"}
	fmt.Println("c4=",*c4)

}