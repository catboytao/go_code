package main

import (
	"fmt"
)
//演示类型断言
type Point struct {
	x int
	y int
}


func main() {
	var a interface{}
	var p Point = Point{1,2}
	a = p
	//b = a  不可以
	b,flag := a.(Point)  //类型断言(带检测的)
	if flag {
		fmt.Println("convert successful")
		fmt.Println(b)
	}else {
		fmt.Println("convert fail")
	}

}