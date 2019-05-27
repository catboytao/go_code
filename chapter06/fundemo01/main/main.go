package main
import (
	"fmt"
	util "go_code/chapter06/fundemo01/utils" //给包名去别名
)



func main(){
	var n1 float64 = 4.5
	var n2 float64 = 3.3
	var operator byte = '-'
	result := util.Cal(n1,n2,operator)
	fmt.Println("result=",result)
}