package main
import (
	"fmt"
)
func main(){
	// 练习1
	var max int = 100
	var count int
	var sum int
	for i := 1; i <= max; i++{
		if i % 9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("count=%d,sum=%d\n",count,sum)

	//练习2
	var n int = 6
	for i := 0; i <= n; i++ {
		fmt.Printf("%d + %d = %d\n",i,(n-i),n)
	}

}