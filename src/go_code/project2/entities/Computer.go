package entities

import "fmt"

type Computer struct {
}

func (computer *Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func ComputerTest1() {
	fmt.Println("Wu Fuqiang")
	computer := Computer{}
	phone := &Phone{}
	computer.Working(phone)
}

func ComputerTest2() {
	fmt.Println("测试类型的向下转换###############")
	var usb Usb
	var phone *Phone = &Phone{}
	usb = phone

	usb.Start()

	phone2, turnSuccess := usb.(*Phone)
	if !turnSuccess {
		fmt.Println("类型转换失败usb.(*Phone)")
		return
	}
	phone2.Study()

	if phone3, flag := usb.(*Phone); flag {
		phone3.Study()
	}
}
