package main

import (
	"fmt"
	"flag"
)

func main() {
	//定义几个变量，用来接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int
	flag.StringVar(&user,"u","","用户名，默认为空")
	flag.StringVar(&pwd,"pwd","","密码，默认为空")
	flag.StringVar(&host,"h","localhost","主机名，默认为localhost")
	flag.IntVar(&port,"port",3306,"端口号，默认为3306")
	//这是一个非常重要的操作，转换，必须调用该方法
	flag.Parse()
	//数据结果
	fmt.Printf("user=%v pwd=%v host=%v port=%v",
		user,pwd,host,port)

}