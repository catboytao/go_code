package main
import (
	"fmt"
)

func main(){
	//map的定义
	/*
		map的key是不能重复的
		map的key-value是无序的
	*/
	//方式1
	//在使用map前，需要先make,make的作用就是给map分配数据空间
	var a map[string]string 
	a = make(map[string]string, 3)
	a["no1"] = "宋江"
	fmt.Println(a)

	//方式2
	cities := make(map[string]string,3)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	//方式3
	heroes := map[string]string{
		"hero1":"宋江",
		"hero2":"卢俊义",
	}
	fmt.Println(heroes)
	
	//定义一个复杂的map
	studentMap := make(map[string]map[string]string) 
	studentMap["stu01"] = make(map[string]string)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京"

	studentMap["stu02"] = make(map[string]string)
	studentMap["stu02"]["name"] = "marry"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "上海"
	fmt.Println(studentMap)

}