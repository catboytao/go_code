package main
import "fmt"

//定义全局变量
var n1 = 100
var n2 = 200
var name = "jack"
//等同于
var (
	n3 = 100
	n4 = 300
	name2 = "marry"
)
	

func main(){
	fmt.Println("n1 =",n1,"n2 =",n2,"name =",name)
	fmt.Println("n3 =",n3,"n4 =",n4,"name2 =",name2)
}