package main
import (
	"fmt"
)

func main(){
	//map的增删改查
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	//map的更新
	cities["no3"] = "上海！"
	fmt.Println(cities)
	
	//map的删除
	delete(cities,"no1")
	fmt.Println(cities)
	//当delete指定的key不存在是，删除不会操作，也不会报错
	delete(cities,"no4")
	fmt.Println(cities)

	//map查找
	val,findRes := cities["no2"]
	if findRes {
		fmt.Printf("有no1 key 值为%v\n",val)
	}else {
		fmt.Printf("没有no1 key\n")
	}

}