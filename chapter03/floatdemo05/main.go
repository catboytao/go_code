package main
import (
	"fmt"
)
//演示golang中的小数类型使用
//小数类型分类 单精度float32  双精度float64 (都是有符号的)
func main(){
	var price float32 = 89.12
	fmt.Println("price=",price)
	var num1 float32 = -0.3433532
	var num2 float64 = -3432353
	fmt.Println("num1=",num1,"num2=",num2)

	//浮点型使用细节
	//golang的浮点类型默认为float64
	var num3 = 1.1
	fmt.Printf("num3的类型是%T\n",num3)

	//浮点型常量有两种表现形式
	// 1.十进制形式
	num4 := 5.12
	num5 := .123 //=>0.123
	fmt.Println("num4=",num4,"num5=",num5)
	//2.科学计数法形式
	num6 := 5.1234e2  //=> 5.1234 * 10的2次方
	fmt.Println("num6=",num6)

	
}