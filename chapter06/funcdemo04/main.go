package main
import (
	"fmt"
)

//在Go中，函数也是一种数据类型
//可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过该变量可以对函数调用

func getSum(n1 int,n2 int) int {
	return n1 + n2
}
//给func(int,int) int 取名为myFunType
type myFunType func(int,int) int

//函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用
func myFun(funvar myFunType,num1 int,num2 int) int {
	return funvar(num1,num2)
}

// go中支持对函数返回值命名
func getSumAndSub(n1 int,n2 int) (sum int,sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}
//交换两个数的值
func swap(n1 *int,n2 *int) (int,int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
	return *n1,*n2
}

//案例演示：编写一个函数sum,可以求出 1到多个int的和
// 定义可变参数，args是slice 切片,可变参数需要放到形参列表最后
func sum(n1 int,args... int) int {
	sum := n1
	// 遍历args
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}


func main() {
	a := getSum
	fmt.Printf("a的类型是%T,getSum的类型是%T\n",a,getSum)

	res := a(10,20)  // 等价于 res := getSum(10,20)
	fmt.Println("res=",res)

	res2 := myFun(a,20,30)
	fmt.Println("res2=",res2)

	// 给int取别名，在go中myInt和int虽然都是int类型，但是go认为myInt和int是两种类型
	// 案例1
	type myInt int
	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1) // go认为myInt 和 int是两种类型
	fmt.Println("num1=",num1,"num2=",num2)
	//案例2
	res3 := myFun(getSum,10,60)
	fmt.Println("res3=",res3)

	a1,a2 := getSumAndSub(30,10)
	fmt.Println("sum=",a1,"sub=",a2)
	// 测试可变参数调用
	res4 := sum(10,0,3,4,5,6)
	fmt.Println("res4=",res4)


	n1 := 10
	n2 := 20
	n1,n2 = swap(&n1,&n2)
	fmt.Println("n1=",n1,"n2=",n2)
	
}