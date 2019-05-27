package main
import (
	"fmt"
)

//定义一个全局匿名函数
var (
	Fun1 = func (n1 int,n2 int) int {
		return n1 * n2
	}
)

func main(){
	// 匿名函数演示
	res := func (n1 int,n2 int) int {
		return n1 + n2
	}(10,20)
	fmt.Println("res=",res)
	
	// 将匿名函数func (n1 int,n2 int) int赋给a变量
	//则a的数据类型就是函数类型，此时，我们可以通过a完成调用
	a := func (n1 int,n2 int) int {
		return n1 - n2
	}
	res2 := a(20,10)
	fmt.Println("res2=",res2)

	// 全局匿名函数的使用
	res3 := Fun1(3,5)
	fmt.Println("res3=",res3)
}