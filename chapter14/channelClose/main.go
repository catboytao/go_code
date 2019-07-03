package main

import (
	"fmt"
)

//演示管道的关闭，管道关闭后不能在往里面写数据，但可以读

func main() {
	intChan := make(chan int,3)
	intChan<- 100
	intChan<- 200
	//close()内置函数，用来关闭管道
	//close(intChan)
	intChan<- 300
	n1 := <-intChan
	fmt.Println("n1=",n1)

	//遍历管道
	intChan2 := make(chan int,100)
	for i :=0;i < 100; i++ {
		intChan2<- i*2 //放入100个数据到管道
	}
	close(intChan2)
	//遍历管道不能使用普通的for循环，应使用for-range
	//在遍历时，如果channel没有关闭，数据可以全部取出，但是会出现死锁
	for v := range intChan2 {
		fmt.Println("v=",v)
	}

}