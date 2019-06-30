package main

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := &Monster{"牛魔王",21,"牛魔拳"}
	var path = "d:/monster2.txt"
	flag := monster.Store(path)
	if flag == false {
		t.Fatalf("Store()执行失败")
	}
	t.Logf("Store()执行成功")
}
func TestRestore(t *testing.T) {
	monster := &Monster{}
	var path = "d:/monster2.txt"
	flag := monster.Restore(path)
	if flag == false {
		t.Fatalf("Restore()执行失败")
	}
	if monster.Name != "牛魔王" {
		t.Fatalf("Restore()错误，希望为=%v 实际为%v","牛魔王",monster.Name)
	}
	t.Logf("Restore()执行成功")
}
