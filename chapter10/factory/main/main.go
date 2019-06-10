package main

import (
	"fmt"
	"go_code/chapter10/factory/model"
)

func main() {
	//通过工厂模式创建student实例
	var stu = model.NewStudent("tom",97.4)
	fmt.Println(*stu)
	fmt.Println("name=",stu.Name,"score=",stu.GetScore())
}