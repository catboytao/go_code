package main
import (
	"fmt"
	"strings"
)

//累加器
// AddUpper函数返回的是一个匿名函数，这个匿名函数引用到函数外的n，因此这个匿名函数就和n形成一个整体，构成闭包
func AddUpper() func(int) int {
	var n int = 10
	return func (x int) int {
		n = n + x
		return n
	}
}


//闭包最佳实践
func makeSuffix(suffix string) func(string) string {
	return func (fileName string) string {
		if strings.HasSuffix(fileName,suffix) {
			return fileName
		}
		return fileName + suffix
	}
}

func main(){
	//累加器
	f := AddUpper()
	fmt.Println(f(1)) // 11

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("winter"))

}