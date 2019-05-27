package main
import (
	"fmt"
	"strconv"
)
// string转基本数据类型
func main(){
	//string -> bool
	var str string = "true"
	var b bool 
	// strconv.ParseBool 会返回两个值（value bool,err error）
	b , _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n",b,b)
	// string -> int
	var str2 string = "2"
	var n1 int64 
	n1, _ = strconv.ParseInt(str2,10,64)
	fmt.Printf("n1 type %T n1=%v\n",n1,n1)
	// string -> float
	var str3 string = "123.456"
	var n2 float64
	n2,_ = strconv.ParseFloat(str3,64)
	fmt.Printf("n2 type %T n2=%v\n",n2,n2)

}