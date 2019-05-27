package main
import (
	"fmt"
)

// 求斐波那契数
func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	}else {
		return fbn(n-1) + fbn(n-2)
	}


}

func getNum(n int) int {
	if n == 10 {
		return 1
	}else {
		return (getNum(n+1) + 1)*2
	}
}

func main(){
	res := fbn(5)
	fmt.Println("res=",res)
	num := getNum(1)
	fmt.Println("num=",num)
}