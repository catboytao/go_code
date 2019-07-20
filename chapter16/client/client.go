package main
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	//1.客户端发送单行数据，然后退出
	for {
		reader := bufio.NewReader(os.Stdin)
		line,err := reader.ReadString('\n')
		line = strings.Trim(line,"\r\n")
		if line == "exit" {
			break
		}
		if err != nil {
			fmt.Println("ReadString err=",err)
		}else{
			_,err := conn.Write([]byte(line+"\n"))
			if err != nil {
				fmt.Println("conn.Write err=",err)
				return
			}
			//fmt.Printf("客户端发送了%d字节的数据",n)
		}
	}

}