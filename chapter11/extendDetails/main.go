package main

import (
	"fmt"
)


/*
	继承细节：
	1.结构体可以使用嵌套匿名结构体所有的字段和方法，
	即：首字母大写或者小写的字段、方法
	2.当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则，
	如希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分。

*/
type A struct {
	Name string
	age int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk",a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello",a.Name)
}

func (b *B) SayOk() {
	fmt.Println("B SayOk",b.Name)
}
type B struct {
	A
	Name string
}


func main() {
	var b B
	b.A.Name = "tom"
	b.A.age = 20
	b.A.SayOk()
	b.A.hello()
	//上面的写法可以简化
	b.Name = "smith"
	b.age = 21
	b.SayOk()
	b.hello()
}