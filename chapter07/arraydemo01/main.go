package main
import (
	"fmt"
)

func main(){
	//定义一个长度为3的数组
	var intArr [3]int
	intArr[0] = 1
	intArr[1] = 3
	intArr[2] = 5
	fmt.Println(intArr)
	fmt.Printf("intArr的地址=%p\n",&intArr) //intArr的地址就是数组第一个元素的地址,数组前后元素地址相差数据类型所占字节数
	// var score [5]float64
	// for i := 0 ; i < len(score); i++ {
	// 	fmt.Printf("请输入第%d个元素的值",i+1)
	// 	fmt.Scanln(&score[i])
		
	// }
	// fmt.Println(score)

	//四种初始化数组的方式
	var numArr01 [3]int = [3]int{1,2,4}
	fmt.Println("numArr01=",numArr01)

	var numArr02 = [3]int{2,4,5}
	fmt.Println("numArr02=",numArr02)
	// 这里[...]是规定的写法
	var numArr03 = [...]int{2,4,6}
	fmt.Println("numArr03=",numArr03)
	// 指定某个位置的值
	var numArr04 = [...]int{1: 300, 0: 349, 2: 342}
	fmt.Println("numArr04=",numArr04)

	//数组的遍历
	//for-range
	for index,value := range numArr01 {
		fmt.Printf("numArr01[%v]=%v\t",index,value)
	}
}