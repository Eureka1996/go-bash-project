package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_base/src/go_code/project2/demo"
	"go_base/src/go_code/project2/entities"
	"go_base/src/go_code/project2/mysort"
	"go_base/src/go_code/project2/utils/arrayUtils"
	"go_base/src/go_code/project2/utils/fileUtils"
	"go_base/src/go_code/project2/utils/flagUtils"
	"os"
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

	fmt.Println("wufuqiang---len(os.Args):", len(os.Args))
	fmt.Println("--------命令行参数--------------")
	for _, value := range os.Args {
		fmt.Println("命令行参数：", value)
	}

	flagUtils.ParseArgs()

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

	// 接口实现
	fmt.Println("----------接口实现--------------")
	entities.ComputerTest1()
	entities.ComputerTest2()

	fmt.Println("--------Struct排序-------")
	// struct排序
	mysort.SortIntTest()
	mysort.SortStructTest()

	fmt.Println("----------判断参数类型-----------")
	demo.TypeJudge(32, 32.32, "32", entities.Phone{Name: "vivo"}, &entities.Phone{Name: "华为"})

	fmt.Println("----------文件操作--------------")

	fileName := "src/resource/properties.conf"
	fileUtils.GetConfig(fileName)
	fileUtils.GetConfigOneTime(fileName)
	fileName2 := "src/resource/wufuqiang.conf"
	fileUtils.WriteFile(fileName2)
	exist, err := fileUtils.PathExist(fileName2)
	fmt.Println("\n", exist, err)

	readFilePath := "src/resource/R-C.jpeg"
	writeFilePath := "src/resource/R-C_copy.jpeg"

	fileUtils.MyCopy(readFilePath, writeFilePath)

}
