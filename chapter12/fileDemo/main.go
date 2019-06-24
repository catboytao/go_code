package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("d:/itcast.log")
	if err != nil {
		fmt.Println("open file err=",err)
	}
	fmt.Printf("file=%v",file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=",err)
	}
}

