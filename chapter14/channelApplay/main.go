package main

import (
	"fmt"
	_"time"
)

//案例：
//1.定义writeData的协程，往channel中写入50个数据
//2.定义一个readData的协程，读取channel中的数据

//思路：
//定义两个channel，一个用来放数据，另一个用来结束主线程使用
//定义一个写数据的方法，写完50个数据后关闭管道，然后定义一个读数据的方法
//读完数据后向另一个管道写入一个bool值，主函数循环读取管道中的值，当读到这个值后，结束循环，程序结束

func WriteData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan<- i
		fmt.Printf("writeData 写入数据=%v\n",i)
		//time.Sleep(time.Second)
	}
	close(intChan)
}

func ReadData(intChan chan int,exitChan chan bool) {
	for {
		k,ok := <-intChan 
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n",k)
	}
	exitChan<- true
	close(exitChan)
}


func main() {
	intChan := make(chan int,50)
	exitChan := make(chan bool,1)
	go WriteData(intChan)
	go ReadData(intChan,exitChan)
	//time.Sleep(time.Second*10)
	for {
		_,ok := <-exitChan
		if !ok {
			break
		}
	}
	
}