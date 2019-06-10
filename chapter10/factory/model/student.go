package model

//定义一个结构体

type student struct {
	Name string
	score float64
}

//因为student结构体首字母是小写，因此只能在model包中使用
//我们通过工厂模式来解决

func NewStudent(n string,s float64) *student {
	return &student{
		Name : n,
		score : s,
	}
}

func (stu *student) GetScore() float64 {
	return stu.score
}