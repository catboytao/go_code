package model
import (
	"fmt"
)

/*
	封装的演示
*/
type person struct {
	Name string
	age int
	salary float64 
}

func NewPerson(name string) *person {
	return &person{
		Name:name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 || age < 150 {
		p.age = age
	}else {
		fmt.Println("年龄范围不正确...")
	}
}

func (p *person) SetSalary(salary float64) {
	p.salary = salary
}

func (p *person) GetAge() int {
	return p.age
}
func (p *person) GetSalary() float64 {
	return p.salary
}