package main
import (
	"fmt"
	//为了使用utils.go文件的变量或者函数，我们需要先引入该model包
	"go_code/chapter03/demo01/model"
)



func main(){
	//定义变量（声明变量）
	var i int
	//给i赋值
	i = 10
	//使用变量
	fmt.Println("i=",i)
	//使用utils.go文件的变量，注：变量名开头大写（公有），变量名开头小写（私有）
	fmt.Println(model.StuName)
}