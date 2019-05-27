package main
import (
	"fmt"
) 
//演示golang中基本数据类型的转换
func main(){
	var i int32 = 100
	var j float32 = float32(i)
	var k int8 = int8(i)
	fmt.Printf("i=%v j=%v k=%v\n",i,j,k)
	//被转换的是变量存储的数据（即值）,变量本身的数据类型并没有变化
	fmt.Printf("i的类型是%T\n",i)

	var num1 int64 = 99999
	var num2 int8 = int8(num1)
	fmt.Println("num2=",num2)

}