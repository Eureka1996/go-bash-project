package main

import "fmt"

func test(n1 int, n2 int, operator string) float32 {
	fmt.Printf("算数计算：%d %s %d = ?\n", n1, operator, n2)
	var result float32
	switch operator {
	case "+":
		result = float32(n1 + n2)
	case "-":
		result = float32(n1 - n2)
	case "*":
		result = float32(n1 * n2)
	case "/":
		result = float32(n1 * 1.0 / n2)
	}
	return result
}

func main() {
	fmt.Println("wufuqiang")
	result := test(12, 45, "*")
	fmt.Println(result)
}
