package main

import (
	"fmt"
	"go_code/chapter11/encapsulate/model"
)

func main() {
	p := model.NewPerson("Tom")
	p.SetAge(20)
	p.SetSalary(5000.0)
	fmt.Println(p.Name,"age=",p.GetAge,"salary=",p.GetSalary)
}