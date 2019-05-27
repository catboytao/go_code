package main
import (
	"fmt"
)
func main(){
	//golang中没有while/do while，可以使用for实现
	// while实现
	var i int = 1
	for {
		if i > 10{
			break
		}
		fmt.Println("hello world!")
		i++
	}
	fmt.Println("i=",i)

	// do while实现
	var j int = 1
	for {
		fmt.Println("hello ok")
		j++
		if i > 10 {
			break
		}
	}
}