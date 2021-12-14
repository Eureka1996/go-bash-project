package main

import (
	"fmt"
	"go-bash-project/src/go_code/project1/pointType"
	"unsafe"
)

func main() {
	// golang程序中整型变量在使用时，遵守保小不保大的原则
	// 即：在保证程序正确运行下，尽量使用占用空间小的数据类型
	var n1 = 1000
	fmt.Printf(" n1的数据类型：%T，\n n1占用的字节数是：%d \n", n1, unsafe.Sizeof(n1))
	var n2 int16 = 100
	fmt.Printf(" n2的数据类型：%T，\n n2占用的字节数是：%d \n", n2, unsafe.Sizeof(n2))

	var name string = "wufuqiang"
	var age int = 20
	var price float64 = 23.5525
	var married bool = false
	var info string = fmt.Sprintf("名称：%s，年龄：%d，价格：%f，是否结婚：%t", name, age, price, married)
	fmt.Println(info)

	fmt.Println("heroName:" + pointType.HeroName)
}
