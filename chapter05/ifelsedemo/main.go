package main
import (
	"fmt"
)

func main(){
	// if age := 20;age > 18 {
	// 	fmt.Println("你的年龄大于18，要对自己负责！")
	// }else {
	// 	fmt.Println("你的年龄不大，这次放过你！")
	// }

	var score int
	fmt.Println("请输入成绩:")
	fmt.Scanln(&score)
	//多分支判断
	if score == 100{
		fmt.Println("奖励一辆BMW")
	}else if score > 80 && score <= 90{
		fmt.Println("奖励一台iphone7")
	}else if score >= 60 && score <= 80{
		fmt.Println("奖励一个 iPad")
	}else{
		fmt.Println("什么都不奖励！")
	}


}