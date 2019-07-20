package main

import (
	"fmt"
)

func main(){

	//1.在默认情况下，管道是双向的（可读可写）
	var chan1 chan int
	chan1 = make(chan int,10)

	//2.声明为只写
	var chan2 chan<- int
	chan2 = make(chan int,20)

	//3.声明为只读
	var chan3 <-chan int
	chan3 = make(chan int,10)

	fmt.Println("chan1=",chan1)
	fmt.Println("chan2=",chan2)
	fmt.Println("chan3=",chan3)
}