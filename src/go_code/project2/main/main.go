package main

import (
	"fmt"
)

func errorHandle() {
	num1 := 10
	num2 := 0
	fmt.Println(num1 / num2)
}

func main() {
	fmt.Println("wufuqiang")
	errorHandle()
	fmt.Println("maoyujiao")
}
