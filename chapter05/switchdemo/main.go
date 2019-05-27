package main
import (
	"fmt"
)

func main(){
	// case/switch后是一个表达式(即：常量值、变量、一个有返回值的函数等)
	// case和switch后的值类型要一致
	// var day byte 
	// fmt.Println("请输入一个字符 a,b,c,d,e,f,g")
	// fmt.Scanf("%c",&day)
	// switch day {
	// 	case 'a':
	// 		fmt.Println("周一")
	// 	case 'b':
	// 		fmt.Println("周二")
	// 	case 'c':
	// 		fmt.Println("周三")
	// 	default:
	// 		fmt.Println("输入有误...")
	// 	}

	// switch 的穿透 fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough  //默认只能穿透一层 
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到")
	}

}