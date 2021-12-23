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
