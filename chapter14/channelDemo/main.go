package main

import (
	"fmt"
)

func main() {
	//演示一下管道的使用
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int,3)
	fmt.Println("intChan=",intChan)
	//2.向管道写入数据
	intChan<-10
	num := 211
	intChan<- num
	intChan<-25
	//注意：当我们给管道写入数据时，不能操作它的容量
	fmt.Printf("channel len=%v cap=%v\n",len(intChan),cap(intChan))

	//3.从管道中取出数据
	//注意：在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报死锁。

	n1 := <-intChan
	fmt.Println(n1)
	n2 := <-intChan
	n3 := <-intChan
	fmt.Println("n2=",n2,"n3=",n3)
	n4 := <-intChan
	fmt.Println(n4)
}