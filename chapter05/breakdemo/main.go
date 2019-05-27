package main
import (
	"fmt"
	//"math/rand"
	//"time"
)
func main(){
	// 我们为了生成一个随机数，还需要给rand设置一个种子
	//time.Now().Unix() :返回一个从1970:01:01的0时0分0秒到现在的秒数
	// rand.Seed(time.Now().Unix())
	// 如何随机的生成1-100整数
	// n := rand.Intn(100) + 1
	// fmt.Println(n)	 

	// var count int
	// for {
	// 	rand.Seed(time.Now().UnixNano())
	// 	n := rand.Intn(100) + 1
	// 	//fmt.Println(n)
	// 	count++
	// 	if n == 99 {
	// 		break
	// 	}
	// }
	// fmt.Println("生成 99 一共使用了 ",count)

	// 指定标签的形式使用break
	// lable1:
	// for i := 0; i < 4; i++ {
	// 	for j := 0;j < 10; j++ {
	// 		if j == 2 {
	// 			break lable1  //跳出lable1
	// 		}
	// 		fmt.Println("j=",j)
	// 	}
	// }

	// break练习1
	// var sum int
	// for i := 1; i <= 100; i++ {
	// 	sum += i
	// 	if sum > 20 {
	// 		fmt.Println("当sum>20时，当前数是：",i)
	// 		break
	// 	}
	// }
	// 练习2
	var username string
	var password string
	var count int = 4
	for i := 1; i <= count; i++ {
		fmt.Println("请输入用户名")
		fmt.Scanln(&username)
		fmt.Println("请输入密码")
		fmt.Scanln(&password)
		if username == "张无忌" && password == "888"{
			fmt.Println("登录成功！")
			break
		}else {
			count--
			fmt.Printf("你还有%d次机会,请珍惜\n",count)
		}
		if count == 0 {
			fmt.Println("机会用完，登录失败")
		}
	}
	

}