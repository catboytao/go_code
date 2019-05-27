package main
import (
	"fmt"
)

func main(){
	//演示切片的基本使用
	var intArr [5]int = [...]int{1,2,4,5,6}
	//定义一个切片，方式1
	slice := intArr[1:3] //表示slice引用intArr数组的起始下标为1，最后的下标为3（但不包含3）
	fmt.Println("intArr=",intArr)
	fmt.Println("slice=",slice)
	fmt.Println("slice 的元素个数=",len(slice))
	fmt.Println("slice 的容量 =",cap(slice))
	fmt.Println()
	//方式2
	var slice2 []int = make([]int, 4,10)
	slice2[0] = 10
	slice2[1] = 20
	fmt.Println(slice2)
	fmt.Println("slice2的size=",len(slice2))
	fmt.Println("slice2的cap=",cap(slice2))
	fmt.Println()
	//方式3
	var slice3 []int = []int{1,3,5}
	fmt.Println(slice3)
	fmt.Println("slice3的size=",len(slice3))
	fmt.Println("slice3的cap=",cap(slice3))
}