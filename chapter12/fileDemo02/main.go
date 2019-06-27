package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

func main() {
	file,err := os.Open("d:/hello.txt")
	if err != nil {
		fmt.Println("open file err=",err)
	}
	//当函数退出时，要及时的关闭file
	defer file.Close()

	//创建一个*Reader,是带缓冲的
	/*
	const (
		defaultBufSize = 4096  //默认的缓冲区为4096个字节
	)
	*/
	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF {
			break
		}
	}
	fmt.Println()
	fmt.Println("文件读取结束！")
}