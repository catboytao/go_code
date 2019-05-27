package main
import (
	"fmt"
)
func main(){

	//跳转控制语句continue的使用
	//lable1:
	// for i := 0; i < 3; i++ {
	// 	label2:
	// 	for j := 0;j < 10; j++ {
	// 		if j == 2 {
	// 			continue label2 
	// 		}
	// 		fmt.Println("j=",j)
	// 	}
	// }
	
	// continue练习1 (输出0-100的奇数)
	// for i := 0;i <= 100; i++ {
	// 	if i % 2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	//continue练习2 （统计正数和负数的个数）

	// var positiveCount int 
	// var negativeCount int
	// var num int
	// for {
	// 	fmt.Println("请输入一个整数:")
	// 	fmt.Scanln(&num)
	// 	if num == 0 {
	// 		break
	// 	}
	// 	if num > 0 {
	// 		positiveCount++
	// 		continue
	// 	}
	// 	negativeCount++

	// }
	// fmt.Printf("正数的个数为%v,负数的个数为%v",positiveCount,negativeCount)

	var totalMoney float64 = 100000
	var count int
	for {
		if totalMoney > 50000 {
			totalMoney -= totalMoney * 0.05
			count++
			continue
		}
		totalMoney -= 1000
		count++
		if totalMoney < 1000 {
			break
		}
	}
	fmt.Printf("该人可以经过%v次路口",count)

}