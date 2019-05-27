package main
import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	str := "hello北"
	fmt.Println("str len=",len(str)) //8

	//字符串遍历
	str2 := "hello世界"
	r := []rune(str2)
	for i := 0 ; i < len(r); i++ {
		fmt.Printf("字符=%c\n",r[i])
	}

	//字符串转整数
	n, err := strconv.Atoi("123")
	if	err != nil {
		fmt.Println("转换错误",err)
	}else {
		fmt.Println("转换的结果是",n)
	}
	//整数转字符串
	str = strconv.Itoa(123)
	fmt.Printf("str=%v,str的类型=%T\n",str,str)

	//字符串转 []byte
	var bytes = []byte("hello")
	fmt.Printf("bytes=%v\n",bytes)

	//[]byte转字符串
	str = string([]byte{97,98,99})
	fmt.Printf("str=%v\n",str)

	// 判断某个字符串是否包含某个字符串
	b := strings.Contains("seafood","foo")
	fmt.Printf("b=%v\n",b)
	// 判断两个字符串是否相等，不区分大小写
	b = strings.EqualFold("abc","Abc")
	fmt.Printf("b=%v\n",b)

	//返回子字符串在字符串第一次出现的index值，如果没有返回-1
	index := strings.Index("NLT_abc","abc")
	fmt.Printf("index=%v\n",index)

	//返回子字符串在字符串最后出现的index值，如果没有返回-1
	index = strings.LastIndex("go golang","go")
	fmt.Printf("index=%v\n",index)
	
	//字符串替换 
	str = strings.Replace("go go hello","go","go语言",-1) // -1 表示全部替换
	fmt.Printf("str=%v\n",str)

	//字符串安装某个字符分割
	strArr := strings.Split("hello,world",",")
	fmt.Printf("strArr=%v\n",strArr)

	//去掉字符串两边的空格
	str = strings.TrimSpace(" tn a lone go pher ntn     ")
	fmt.Printf("str=%v\n",str)

	//去掉两边指定的字符
	str = strings.Trim(" ! hello!"," !")
	fmt.Printf("str=%q\n",str)

	//判断字符串是否以指定的字符串开头
	b = strings.HasPrefix("http://www.baidu.com","http")
	fmt.Printf("b=%v",b)
	//判断字符串是否以指定的字符串结尾
	b = strings.HasSuffix("hello.jpg",".jpg")
	fmt.Printf("b=%v",b)
}