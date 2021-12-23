package entities

import "fmt"

type Phone struct {
}

func (phone *Phone) Start() {
	fmt.Println("手机开始工作... ...")
}

func (phone *Phone) Stop() {
	fmt.Println("手机停止工作... ...")
}
