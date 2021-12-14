// 表示该helloworld.go文件所在的包是main，在go中，每个文件都必须归属于一个包。
package main

import "fmt"

// func是一个关键字，表示一个函数
// main是函数名，是一个主函数，即我们程序的入口
func main() {
	// 调用fmt包的函数Println输出
	fmt.Println("Hello world,Hello WuFuqiang")
}
