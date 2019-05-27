package main
import (
	"fmt"
)
func main(){
	var i int = 5
	//二进制输出
	fmt.Printf("%b",i)
	//八进制：0-7，满8进1，以数字0开头表示
	var j int = 011
	fmt.Println("j=",j)

	//十六进制，以0x或0X表示
	var k int = 0x11
	fmt.Println("k=",k)
}