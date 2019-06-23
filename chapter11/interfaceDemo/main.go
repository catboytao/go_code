package main

import (
	"fmt"
)
//定义一个Usb接口，有两个方法
type Usb interface {
	Start()
	Stop()
}
//定义一个Phone结构体
type Phone struct {
	
}
//定义一个Camera结构体
type Camera struct {

}
//Phone结构体实现Usb接口的Start()和Stop()方法
func (p Phone) Start() {
	fmt.Println("手机开始工作.....")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作......")
}

//Camera结构体实现Usb接口的Start()和Stop()方法
func (c Camera) Start() {
	fmt.Println("相机开始工作......")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作......")
}
//定义一个Compute结构体
type Compute struct {

}
//编写一个Working方法，接收一个Usb接口类型变量
//只要实现了Usb接口(所谓实现Usb接口，就是指实现了Usb接口声明的方法)，就能作为参数传递给该方法
func (c Compute) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	compute := Compute{}	
	phone := Phone{}
	camera := Camera{}
	//体现了面向接口的多态特征
	compute.Working(phone)
	compute.Working(camera)

	//定义一个多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{}
	usbArr[1] = Phone{}
	usbArr[2] = Camera{}
	fmt.Println(usbArr)
}
