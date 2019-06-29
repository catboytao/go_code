package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

//演示文件的拷贝

func CopyFile(dstFileName,SrcFileName string) (written int64,err error){
	srcFile, err := os.Open(SrcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
	}

	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)
	
	dstFile ,err := os.OpenFile(dstFileName,os.O_WRONLY | os.O_CREATE,0666)
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)
	return io.Copy(writer,reader)

}

func main() {

	srcFile := "d:/boy.jpg"
	dstFile := "e:/abc.jpg"
	_, err := CopyFile(dstFile,srcFile)
	if err == nil {
		fmt.Println("拷贝完成！")
	}else {
		fmt.Println("拷贝错误 err = ",err)
	}
	


}