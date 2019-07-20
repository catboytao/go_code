package main

import (
	"fmt"
	"net" //做网络socket开发时，net包含了我们需要的所有的方法和函数
)


func process(conn net.Conn) {
	//这里我们循环的接收客户端发送的数据
	defer conn.Close()
	for {
		//创建一个新的切片
		buf := make([]byte,1024)
		//fmt.Printf("服务器在等待客户端%s 发送信息\n",conn.RemoteAddr().String())
		n,err := conn.Read(buf)  //等待客户端通过conn发送信息，如果客户端没有writer[发送]，那么协程就阻塞
		if err != nil {
			fmt.Println("服务器端读取错误 err=",err)
			return
		}
		//显示客户端发送的内容到服务器终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {

	fmt.Println("服务器开始监听.....")
	listen,err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=",err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端来链接....")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=",err)
		}else {
			fmt.Printf("conn=%v 客户端ip=%v\n",conn,conn.RemoteAddr().String())
		}
		go process(conn)
	}
	//fmt.Println("listen=",listen)
}
