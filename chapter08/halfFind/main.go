package main
import (
	"fmt"
)

func find(n int,arr [5]int) {
	var start int = 0
	var end int = len(arr) - 1
	var mid int 
	for {
		mid = (end + start) / 2
		if start > end {
			fmt.Println("没有找到！")
			break
		}
		if n > arr[mid] {
			fmt.Println("猜小了！")
			start = mid + 1
		}else if n < arr[mid] {
			fmt.Println("猜大了！")
			end = mid - 1
		}else {
			fmt.Println("该数字对应的索引为：",mid)
			break
		}

	}
}


func main(){
	var arr [5]int = [...]int{1,2,3,4,5}
	find(6,arr)
}