package main 

import (
	"fmt"
)

type MethodUtils struct {

}

func (m MethodUtils) Print(n1 int,n2 int) {
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var mu MethodUtils
	mu.Print(6,6)
}