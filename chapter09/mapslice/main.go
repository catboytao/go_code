package main

import (
	"fmt"
)

func main() {
	//定义一个map切片
	var monster []map[string]string
	monster = make([]map[string]string,2)
	if  monster[0] == nil {
		monster[0] = make(map[string]string,2)
		monster[0]["name"] = "牛魔王"
		monster[0]["age"] = "500"
	}

	if monster[1] == nil {
		monster[1] = make(map[string]string,2)
		monster[1]["name"] = "玉兔精"
		monster[1]["age"] = "400"
	}

	newMonster := map[string]string{
		"name" : "火云邪神",
		"age" : "200",
	}
	monster = append(monster,newMonster)
	fmt.Println(monster)

}