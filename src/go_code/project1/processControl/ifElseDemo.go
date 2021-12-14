package main

import "fmt"

func main() {

	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanf("%d", &age)
	if age > 18 {
		fmt.Println("已成年")
	} else {
		fmt.Println("未成年")
	}
	fmt.Println("----------------")

	fmt.Println("请再次输入年龄：")
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Println("未成年")
	} else if age < 30 {
		fmt.Println("还未到而立之年")
	} else if age < 40 {
		fmt.Println("还未到不惑之年")
	} else {
		fmt.Println("年纪太大了")
	}

	if name := "wfq"; name == "wfq" {
		fmt.Printf("姓名为：%v \n", name)
	}
}
