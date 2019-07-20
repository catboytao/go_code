package main
import (
	"fmt"
	"reflect"
)

//专门演示反射
func reflectTest(b interface{}){
	//通过反射获取传入的变量的type,kind,值
	//1.获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=",rType)
	//2.获取reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := rVal.Int() + 2
	fmt.Println(n2)
	fmt.Printf("rVal=%v type=%T\n",rVal,rVal)

	//将rVal 转换成 interface{}
	iV := rVal.Interface()
	//将interface{}通过类型断言装换成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=",num2)
}

//演示对结构体的反射

type Student struct {
	Name string
	Age int
}

func reflectTest02(b interface{}) {
	//1.获取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=",rType)
	//2.获取reflect.Value
	rVal := reflect.ValueOf(b)
	//3.获取变量对应的Kind(类别)
	kind1 := rVal.Kind()
	kind2 := rType.Kind()
	fmt.Printf("kind1=%v kind2=%v\n",kind1,kind2)
	iV := rVal.Interface()
	fmt.Printf("iV=%v type=%T\n",iV,iV)
	stu,ok := iV.(Student)
	if ok{
		fmt.Println("stu.Name=",stu.Name)
	}
}

func main() {
	//var num int = 100
	//reflectTest(num)
	stu := Student{
		Name:"tom",
		Age:20,
	}
	reflectTest02(stu)
}
