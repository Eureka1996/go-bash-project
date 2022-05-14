package main

import "fmt"

/**
 * @author wufuqiang
 * @date 2022/5/15 00:16
 */

// 初始化顺序：全局变量初始化 -> init() -> main()

var i int = initVar()

func init() {
	fmt.Println("init-----")
}

func initVar() int {
	fmt.Println("initVar------")
	return 111
}

func main() {

	fmt.Println("main------")

}
