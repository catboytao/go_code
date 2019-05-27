package main
import (
	"fmt"
)

func main(){
	//冒泡排序
	var arr [5]int = [...]int{4,3,6,2,5}
	var temp int
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - 1 - i; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println("冒泡排序过后的arr=",arr)


}