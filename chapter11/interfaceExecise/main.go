package main
import (
	"fmt"
	"math/rand"
	"sort"
)

/*
	接口的应用：对一个结构体切片排序
*/

//1.定义一个Hero结构体
type Hero struct {
	Name string
	Age int
}

//2.定义一个Hero结构体切片
type HeroSlice []Hero

//3.实现Interface接口
func (h HeroSlice) Len() int {
	return len(h)
}
// Less方法就是决定你使用什么标准进行排序
// 1.按Hero的年龄从小到大排序
func (h HeroSlice) Less(i,j int) bool {
	return h[i].Age < h[j].Age  
}
func (h HeroSlice) Swap(i,j int) {
	tmp := h[i]
	h[i] = h[j]
	h[j] = tmp
}
func main() {
	//测试对结构体切片进行排序
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name : fmt.Sprintf("英雄~%d",rand.Intn(100)),
			Age : rand.Intn(100),
		}
		heroes  = append(heroes,hero)
	}
	//查看排序前的顺序
	for _,v := range heroes {
		fmt.Println(v)
	}
	//调用sort.Sort
	sort.Sort(heroes)
	//查看排序后的顺序
	fmt.Println("-------排序后-------")
	for _,v := range heroes {
		fmt.Println(v)
	}
}