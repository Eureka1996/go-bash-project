package entities

import "fmt"

type Phone struct {
	Name string
}

func (phone *Phone) Start() {
	fmt.Println("手机开始工作... ...")
}

func (phone *Phone) Stop() {
	fmt.Println("手机停止工作... ...")
}

func (phone *Phone) Study() {
	fmt.Println("手机可用来学习")
}
