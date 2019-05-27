package main
import (
	"fmt"
	"strconv"
)
//基本类型转string类型

func main(){
	
	//方式1
	var a int = 99
	var b float64 = 23.435
	var c bool = true
	var d byte = 'h'
	var str string
	// int转string
	str = fmt.Sprintf("%d",a)
	fmt.Printf("str type %T str=%q\n",str,str)
	//float 转string
	str = fmt.Sprintf("%f",b)
	fmt.Printf("str type %T str=%q\n",str,str)
	//bool 转string
	str = fmt.Sprintf("%t",c)
	fmt.Printf("str type %T str=%q\n",str,str)
	//byte转string
	str = fmt.Sprintf("%c",d)
	fmt.Printf("str type %T str=%q\n",str,str)
	fmt.Println("------------")
	//第二种方式 strconv包
	var e int = 99
	var f float64 = 34.543
	var g bool = true
	//int -> string
	str = strconv.FormatInt(int64(e),10)
	//str = strconv.Itoa(e)
	fmt.Printf("str type %T str=%q\n",str,str)
	//float -> string
	str = strconv.FormatFloat(f,'f',10,64)
	fmt.Printf("str type %T str=%q\n",str,str)
	//bool -> string
	str = strconv.FormatBool(g)
	fmt.Printf("str type %T str=%q\n",str,str)
}