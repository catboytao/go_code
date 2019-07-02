package main

import (
	"fmt"
)

type Cat struct {
	Name string 
	age int 
}

func main() {
	//定义一个存放任何类型的管道
	allChan := make(chan interface{},3)
	allChan <- 10
	allChan<- "tom jack"
	cat := Cat{"小花",4}
	allChan<- cat
	
	//我们希望获取到管道中的第三个元素，需要将前2个推出
	<-allChan
	<-allChan

	newCat := <-allChan
	fmt.Printf("newCat=%T newCat=%v",newCat,newCat)
	//下面的写法是错误的，编译时通不过
	//fmt.Println(newCat.Name)
	//使用类型断言
	a := newCat.(Cat)
	fmt.Println(a.Name)
}