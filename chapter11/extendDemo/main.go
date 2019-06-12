package main

import (
	"fmt"
)
/*
	演示继承
*/
//学生结构体
type Student struct {
	Name string
	Age int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名：%v  年龄：%v 成绩：%v",stu.Name,stu.Age,stu.Score)
}
func (stu *Student) SetScore(score int) {
	if score <= 100 && score >= 0 {
		stu.Score = score
	}else {
		fmt.Println("成绩输入有误！")
	}
}

//小学生结构体
type Pupil struct {
	Student
}
//大学生结构体
type Graduate struct {
	Student
}

func (pupil *Pupil) testing() {
	fmt.Println("小学生正在考试...")
}

func (graduate *Graduate) testing() {
	fmt.Println("大学生正在考试....")
}



func main() {
	pupil := Pupil{}
	pupil.Student.Name = "Tom"
	pupil.Student.Age = 19
	pupil.testing()
	pupil.Student.SetScore(80)
	pupil.Student.ShowInfo()
}