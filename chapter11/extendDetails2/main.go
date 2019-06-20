package main

import (
	"fmt"
)

type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func main() {
	var tv = TV {
		Goods {
			Name:"电视机",
			Price:1890,
		},
		Brand {
			Name:"海尔",
			Address:"山东",
		},
	}

	tv2 := TV2 {
		&Goods {
			Name:"电视机002",
			Price:3499,
		},
		&Brand {
			Name:"康佳",
			Address:"北京",
		},
	}
	fmt.Println(tv)
	fmt.Println(*tv2.Goods,*tv2.Brand)

}