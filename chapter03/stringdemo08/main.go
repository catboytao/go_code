package main
import (
	"fmt"
)
//演示golang中string类型使用
//Go语言的字符串的字节使用UTF-8编码标识Unicode文本
// 在Go中字符串是不可变的
func main(){
	var address string = "北京长城 hello world"
	fmt.Println(address)
	
	//使用反引号输出字符串
	str := `
	import (
		"fmt"
		"unsafe"
	)
	//演示golang中bool类型使用
	func main(){
		var b = false
		fmt.Println("b=",b)
		//注意事项
		//1.bool类型占用存储空间是1个字节
		fmt.Printf("b的占用空间为%d字节",unsafe.Sizeof(b))
		//2.bool类型只能取true或者false
		
	}`
	fmt.Println(str)

	//字符串拼接
	var str2 = "hello" + "world"
	str2+= "haha!"
	fmt.Println(str2)

	//各种变量类型的默认值（零值）
	var a int 
	var b float32
	var c float64
	var isMarried bool
	var name string
	//%v，表示按照变量的原始值输出
	fmt.Printf("a=%d,b=%v,c=%v,isMarried=%v,name=%v",a,b,c,isMarried,name)
}