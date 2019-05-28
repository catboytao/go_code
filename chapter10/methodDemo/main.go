package main

import (
	"fmt"
)

type Person struct {
	Name string
}

//给Person类型绑定一个方法
func (p Person) test() {
	fmt.Println("test()",p.Name)
}

func (p Person) speak() {
	fmt.Printf("%v是一个好人\n",p.Name)
}

func (p Person) jisuan() {
	var res int
	for i := 1;i <= 1000; i++ {
		res += i
	}
	fmt.Println("计算的结果为：",res)
}

func (p Person) jisuan2(n int) {
	var res int
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println("计算的结果为：",res)
}

func (p Person) getSum(n1 int,n2 int) int {
	return n1 + n2
}
func main() {
	var p Person = Person{}
	p.Name = "tom"
	//p.test()
	p.speak()
	p.jisuan()
	p.jisuan2(10)
	res := p.getSum(2,4)
	fmt.Println("res=",res)
}