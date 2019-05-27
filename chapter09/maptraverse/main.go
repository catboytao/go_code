package main

import (
	"fmt"
)

func main() {
	// 使用for-range遍历map
	var cities map[string]string
	cities = make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	for k,v := range cities {
		fmt.Printf("key=%v value=%v\n",k,v)
	}
}