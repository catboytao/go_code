package main

import (
	"fmt"
	"bufio"
	"os"
)

//判断文件或者目录是否存在
func IsExists(path string) (bool,error) {
	_,err := os.Stat(path)
	if err == nil {
		return true,nil
	}
	if os.IsNotExist(err) {
		return false,nil
	}
	return false,err
}

func main()  {
	filePath := "d:/abc.txt"
	/*
	const (
    O_RDONLY int = syscall.O_RDONLY // open the file read-only.
    O_WRONLY int = syscall.O_WRONLY // open the file write-only.
    O_RDWR   int = syscall.O_RDWR   // open the file read-write.
    O_APPEND int = syscall.O_APPEND // append data to the file when writing.
    O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
    O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist
    O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
    O_TRUNC  int = syscall.O_TRUNC  // if possible, truncate file when opened.
)
	*/
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