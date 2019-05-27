package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 练习1
	var arr01 [26]byte
	var str byte = 'A'
	for i := 0;i < len(arr01); i++ {
		arr01[i] = str
		str ++
	}
	for i := 0; i < len(arr01); i++ {
		fmt.Printf("%c",arr01[i])
	}
	fmt.Println()
	//练习2
	var intArr = [...]int{1,-1,3,40,23,34}
	var maxNum int
	var maxValIndex int
	for key,value := range intArr {
		if value > maxNum {
			maxNum = value
			maxValIndex = key
		}
	}
	fmt.Println("maxNum=",maxNum,"index=",maxValIndex)

	//练习3
	var intArr2 = [...]int{1,3,4,5,6,7}
	var sum int
	for _,value := range intArr2 {
		sum += value
	}
	var avg float64 = float64(sum) / float64(len(intArr2))
	fmt.Printf("sum=%v,avg=%v\n",sum,avg)

	//练习4
	// 生成一个长度为5的随机数组，并对它进行反转
	var intArr3 [5]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(intArr3); i++ {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println(intArr3)

	var temp int
	for i := 0; i < len(intArr3) / 2; i++ {
		temp = intArr3[i]
		intArr3[i] = intArr3[len(intArr3) - 1 - i]
		intArr3[len(intArr3) - 1 - i] = temp
	}
	fmt.Println(intArr3)

}


