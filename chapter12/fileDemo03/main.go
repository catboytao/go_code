package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content,err := ioutil.ReadFile("d:/hello.txt")
	if err != nil {
		fmt.Println("read file err=",err)
	}
	fmt.Println(content) // []byte
	fmt.Println(string(content))
}