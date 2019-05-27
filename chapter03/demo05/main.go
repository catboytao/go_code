package main
import (
	"fmt"
)

func main(){
	//从控制台接收用户输入
	//方式1.fmt.Scanln
	var name string
	var age byte
	var salary float32
	var isPass bool
	// fmt.Println("请输入姓名:")
	// fmt.Scanln(&name)
	
	// fmt.Println("请输入年龄:")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪水:")
	// fmt.Scanln(&salary)

	// fmt.Println("请输入是否通过考试:")
	// fmt.Scanln(&isPass)

	// fmt.Printf("name=%v\nage=%v\nsalary=%v\nisPass=%v",name,age,salary,isPass)

	//方式2：fmt.Scanf,可以按指定的格式输入
	fmt.Println("请输入你的名字，年龄，薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t",&name,&age,&salary,&isPass)
	fmt.Printf("name=%v\nage=%v\nsalary=%v\nisPass=%v",name,age,salary,isPass)

}