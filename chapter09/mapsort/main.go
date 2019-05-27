package main

import (
	"fmt"
	"sort"
)

func main() {
	//map 的排序
	m := make(map[int]int,10)
	m[10] = 100
	m[1] = 12
	m[4] = 56
	m[8] = 90
	fmt.Println(m)

	//1.先将map的key放入到切片中
	//2.对切片进行排序
	//3.遍历切片，然后按照key来输出map的值

	var keys []int
	for k,_ := range m {
		keys = append(keys,k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _,k := range keys {
		fmt.Printf("m[%v]=%v\n",k,m[k])
	}
 }