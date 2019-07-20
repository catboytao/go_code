package main

import (
	"fmt"
	"time"
)

func main() {

	//通过select解决管道阻塞的问题
	
	//1.定义一个intChan,放10个数据
	var intChan chan int = make(chan int,10)
	for i := 1; i <= 10; i++ {
		intChan<- i
	}

	//2.定义一个stringChan，放5个数据

	var stringChan chan string = make(chan string,5)
	for i := 1; i <= 5; i++ {
		stringChan<- "hello" + fmt.Sprintf("%d",i)
	}
	label :
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读到数据%d\n",v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从stringChan读到数据%s\n",v)
			time.Sleep(time.Second)
		default:
			fmt.Println("都取不到，不玩了！")
			break label
		}
	}
}