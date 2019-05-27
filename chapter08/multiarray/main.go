package main
import (
	"fmt"
)

func main(){
	//二维数组的演示案例
	// var arr [3][4]int 
	// arr[1][2] = 1
	// arr[2][1] = 2
	// arr[2][3] = 3
	// fmt.Println(arr)
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 4; j++ {
	// 		fmt.Print(arr[i][j]," ")
	// 	}
	// 	fmt.Println()
	// }

	// var arr2 [2][3]int = [...][3]int{{1,2,3},{4,5,6}}
	// fmt.Println(arr2)

	// //二维数组的遍历
	// //for 循环来遍历
	// for i := 0; i < len(arr2); i++ {
	// 	for j := 0; j < len(arr2[i]); j++ {
	// 		fmt.Printf("%v\t",arr2[i][j])
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println("--------")
	// //for range来遍历
	// for _,v1 := range arr2 {
	// 	for _,v2 := range v1 {
	// 		fmt.Printf("%v\t",v2)
	// 	}
	// 	fmt.Println()
	// }
	
	//向一个有序的数组插入一个值
	var arr [4]int = [...]int{1,3,5,7}
	var arr2 [5]int
	var num int = 2
	for i := 0; i < len(arr) - 1; i++ {
			
		if num > arr[i] && num < arr[i+1] {
			arr2[i] = arr[i]
			arr2[i+1] = num
			for j := i+1; j < len(arr); j++{
				arr2[j+1] = arr[j]
			}
			break
		}else{
			arr2[i] = arr[i]
		}
	}
	fmt.Println(arr2)
}