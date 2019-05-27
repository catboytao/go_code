package main
import (
	"fmt"
	"go_code/chapter06/funcinit/utils"
)

var num = test()

// 当文件有全局变量的定义时，它会执行在init方法之前
func test() int {
	fmt.Println("test()...")
	return 30
}


//初始化方法，执行在main方法之前
func init(){
	fmt.Println("init().......")
}

func main(){
	fmt.Println("main()......")
	fmt.Println("name=",utils.Name,"age=",utils.Age)
}