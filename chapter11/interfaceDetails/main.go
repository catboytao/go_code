package main

import (
	"fmt"
)

type A interface {
	Say()
}

type Stu struct {
	Name string
}

func (s Stu) Say() {
	fmt.Println("s Say()")
}


func main() {
	var stu Stu
	var a A = stu
	a.Say()
}