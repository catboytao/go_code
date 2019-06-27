package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	filePath := "d:/abc.txt"
	file,err := os.OpenFile(filePath,os.O_WRONLY | os.O_CREATE,0666)
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
		return
	}
	defer file.Close()
	str := "hello,tom\r\n"
	writer := bufio.NewWriter(file)
	for i := 0;i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
	//因为Writer是带缓冲的，因此在调用WriteString方法时，其实内容
	//是先写入到缓存中，所以需要调用Flush方法，将缓冲的数据真正的写入到文件中
	
}