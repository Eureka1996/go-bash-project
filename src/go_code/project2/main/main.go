package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"project2/src/go_code/project2/mysort"
	"project2/src/go_code/project2/utils/arrayUtils"
	"strings"
)

// 异常处理
func errorHandle() {

	defer func() {
		if error := recover(); error != nil {
			fmt.Printf("出现异常，%T \n", error)
			fmt.Println("异常信息：", error)
		}
	}()

	num1 := 10
	num2 := 0
	fmt.Println(num1 / num2)
}

// 自定义异常
func readConfig(fileName string) (err error) {
	if strings.HasSuffix(fileName, "conf") {
		// 读取配置文件
		return nil
	} else {
		return errors.New("读取配置文件错误。")
	}
}

func readConfigTest1() {
	err := readConfig("wufuqiang.cnf")
	if err == nil {
		fmt.Println("读取文件成功。")
	} else {
		// 如果读取文件发生错误，就输出这个错误，并终止程序。
		panic(err)
	}
}

// 数据遍历
func arrayForeachTest() {
	arrInt := [3]int{1, 2, 3}
	arrayUtils.ArrayForeach(arrInt)

	for _, value := range arrInt {
		fmt.Println(value)
	}

	arrayUtils.ArrayForeach2([2][3]int{{4, 5, 6}, {3, 1, 5}})
}

// 结构体
type Aninal struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

func main() {
	fmt.Println("wufuqiang")
	fmt.Println("maoyujiao")

	var aninal Aninal
	aninal.Name = "小动物"
	aninal.Age = 2
	aninal.Color = "大红色"

	fmt.Println(aninal)

	marshal, err := json.Marshal(aninal)

	if err == nil {
		fmt.Println("json字符串：", string(marshal))
	} else {
		fmt.Println(err)
	}

	//entities.ComputerTest1()

	mysort.SortIntTest()
}
