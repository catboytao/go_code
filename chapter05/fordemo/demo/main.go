package main
import (
	"fmt"
)
func main(){
	//for循环快速入门
	// 方式1
	for i := 1; i <= 10; i++ {
		fmt.Println("hello world!")
	}
	//方式2
	j := 1
	for j <= 10 {
		fmt.Println("你好！")
		j++
	}
	//方式3
	k := 1
	for {
		if k <= 10 {
			fmt.Println("ok~")
		}else{
			break
		}
		k++
	}
	// 字符串遍历方式1 传统方式
	// 实际上是按照字节方式遍历，如果字符串有中文，则会乱码，因为在utf8编码中文占3个字节
	// 需要将字符串转换成切片类型 []rune
	var str string = "hello,world重庆"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n",str2[i])
	}
	// 字符串遍历方式2 for-range
	// 实际上是按照字符方式遍历，如果字符串有中卫，也ok
	for index, val := range str {
		fmt.Printf("index=%d,val=%c\n",index,val)
	}


}