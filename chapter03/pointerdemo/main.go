package main
import (
	"fmt"
)
//演示golang中指针类型
func main(){
	//基本数据类型在内存布局
	var i int = 10
	// i的地址是什么，&i
	fmt.Println("i的内存地址为",&i)
	//指针演示
	//1.ptr是一个指针变量
	//2.ptr的类型是 *int
	//3.ptr 本身的值是&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n",ptr)
	fmt.Printf("ptr的内存地址为%v\n",&ptr)
	fmt.Printf("ptr指向的值=%v",*ptr)
}