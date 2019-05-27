package utils
import (
	"fmt"
)

var Name string
var Age int

func init(){
	fmt.Println("utils init().....")
	Name = "卓涛"
	Age = 21
}