package main
import (
	"fmt"
)

func main(){
	//切片的遍历
	//方式1
	var arr [5]int = [...]int{1,3,5,7,9}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n",i,slice[i])
	}
	fmt.Println()
	//方式2
	for k,v := range slice {
		fmt.Printf("slice[%v]=%v\n",k,v)
	}
	fmt.Println()
	//用append内置函数，可以对切片进行动态追加
	var slice2 []int = []int{10,20,30}
	slice2 = append(slice2,40,50)
	//通过append将切片slice2追加给slice2
	slice2 = append(slice2,slice2...)
	fmt.Println("slice2",slice2)

	//切片的拷贝操作
	var slice3 []int = []int{1,2,3,4,5}
	var slice4 = make([]int,10)
	copy(slice4,slice3)
	fmt.Println("slice3=",slice3)
	fmt.Println("slice4=",slice4)

}