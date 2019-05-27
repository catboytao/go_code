package main
import (
	"fmt"
)
func main(){
	// for i := 1; i <= 9; i++{
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%v * %v = %v\t",i,j,i*j)
	// 	}
	// 	fmt.Println()
	// }

	// 统计三个班成绩情况，每个班五个同学，求每个班的平均成绩及所有班级的平均成绩
	var classNum int = 3
	var stuNum int = 5
	var sum float64
	var avg float64
	var passCount int
	for j := 1; j <= classNum; j++{
		var avg_one float64
		var sum_one float64
		var score float64
		for i := 1; i <= stuNum; i++{
			fmt.Printf("请输入第%d个班级第%d个学生的成绩 \n",j,i)
			fmt.Scanln(&score)
			sum_one += score
			if score >= 60{
				passCount++
			}
		}
		avg_one = sum_one / float64(stuNum)
		fmt.Printf("第%v个班级学生的平均成绩为%v\n",j,avg_one)
		sum += sum_one
	}
	avg = sum / float64(classNum * stuNum)
	fmt.Println("所有班级学生的平均成绩为：",avg)
	fmt.Println("及格人数为：",passCount)
	


}