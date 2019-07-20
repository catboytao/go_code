package main

import (
	"fmt"
	"time"
)

func putNum(c chan int) {
	for i := 1;i <= 80000; i++ {
		c<- i
	}
	close(c)
}

func primeNum(intChan chan int,primeChan chan int,exitChan chan bool) {
	var flag bool
	for {
		num,ok := <-intChan
		if !ok  {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan<- num
		}
	}
	fmt.Println("有一个primeNum协程因为取不到数据，退出！")
	exitChan<- true
}

func main() {
	intChan := make(chan int,1000)
	primeChan := make(chan int,20000)
	exitChan := make(chan bool,4)

	start := time.Now().Unix()
	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go primeNum(intChan,primeChan,exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan	
		}
		end := time.Now().Unix()
		fmt.Println("使用协程耗时=",end - start)
		close(primeChan)
	}()

	for {
		_,ok := <-primeChan
		if !ok{
			break
		}
		//fmt.Println(res)
	}

	
}