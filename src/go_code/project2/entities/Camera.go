package entities

import "fmt"

type Camera struct {
}

func (camera *Camera) Start() {
	fmt.Println("相机开始工作... ...")
}

func (camera *Camera) Stop() {
	fmt.Println("相机停止工作... ...")
}
