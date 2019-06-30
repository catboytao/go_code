package main
import (
	"fmt"
	"encoding/json"
)
//struct tag 使用
type Person struct {
	Name string `json:"name"`   //反射机制
	Age int	`json:"age"`
	Skill string `json:"skill"`
}

func main() {
	p := Person{"小明",20,"Python"}
	jsonStr,err := json.Marshal(p)
	if err != nil {
		fmt.Println("json处理有误",err)
	}else {
		fmt.Println("jsonStr=",string(jsonStr))
	}
}