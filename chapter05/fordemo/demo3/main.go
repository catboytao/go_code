package main
import (
	"fmt"
)
func main(){
	
	// for i := 1; i <= 5; i++{
	// 	for j := 1; j <= i; j++{
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
	//打印金字塔（空心）
	var level int = 5
	for i := 1; i <= level; i++{
		for k := 1; k <= level - i; k++{
			fmt.Print(" ")
		}
		//每层第一个和最后一个打印*,其他打印空格
		for j := 1; j <= 2*i -1; j++{
			if j == 1 || j == 2*i -1 || i == level{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}	
		}
		fmt.Println()
	}




}