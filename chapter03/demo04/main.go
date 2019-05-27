package main
import (
	"fmt"
	"unsafe"
)

func main(){
	//int,uint,rune(等价于int32),byte(等价于uint8)的使用
	var a int = 8900
	fmt.Println("a=",a)
	var b uint = 1
	var c byte = 255
	var d rune = 2355
	fmt.Println("b=",b,"c=",c,"d=",d)

	//整数的使用细节
	//fmt.Printf()用于做格式化输出
	var n1 = 100
	fmt.Printf("n1 的类型是%T\n",n1)

	//如何在程序查看某个变量占用字节大小和数据类型
	//unsafe.Sizeof() 可以查看某个变量占用的字节数
	var n2 int64 = 10
	fmt.Printf("n2的数据类型是%T n2占用的字节数是%d",n2,unsafe.Sizeof(n2))

}