package main

import (
	"testing"
)

func TestAddUpper(t *testing.T)  {
	//调用
	res := addUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10)执行错误，期望值=%v 实际值为%v\n",55,res)
		t.Fatalf("AddUpper(10)执行错误，期望值=%v 实际值为%v\n",55,res)
	}
	//如果正确，输出日志
	t.Logf("AddUpper(10) 执行正确...")
}

func TestSub(t *testing.T) {
	res := sub(10,7)
	if res != 3 {
		t.Fatalf("Sub(10,7)执行错误，期望值=%v 实际值为%v\n",3,res)
	}
	t.Logf("Sub(10,7) 执行正确...")
}
