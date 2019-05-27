package main
import "fmt"

func main(){
	//golang的变量使用方式
	//第一种：指定变量类型，声明后若不赋值，使用默认值
	var i int
	fmt.Println("i=",i)

	//第二种：根据值自行判定变量类型（类型推导）
	var num = 10.11
	fmt.Println("num=",num)

	//第三种：省略var,注意 := 左侧的变量不应该是已经声明过的，否则会导致 编译错误
	// name := "tom"
	// fmt.Println("name=",name)

	//golang一次声明多个变量的方式
	//方式1
	var n1,n2,n3 int
	fmt.Println("n1 =",n1,"n2 =",n2,"n3 =",n3)

	//方式2
	// var name, age, sex = "tom", 21, "男"
	// fmt.Println("name =",name,"age =",age,"sex =",sex)

	//方式3
	name,age,sex := "tom", 21, "男"
	fmt.Println("name =",name,"age =",age,"sex =",sex)
}