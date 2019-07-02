package main
import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int)
	//声明一个全局的互斥锁
	lock sync.Mutex
)
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//加锁
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	//开启多个协程完成这个任务
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	//休眠10秒钟
	time.Sleep(10*time.Second)
	//为什么在读的时候还要加锁呢？
	// 因为我们程序设计上可以知道10秒就执行完所有协程，但是主线程并不知道，因此
	//底层可能仍然出现资源争夺，因此加入互斥锁即可解决问题

	//遍历结果
	lock.Lock()
	for i,v := range myMap {
		fmt.Printf("myMap[%v]=%v\n",i,v)
	}
	lock.Unlock()


}
