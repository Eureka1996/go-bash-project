package main

import (
	"fmt"
	"go_base/src/go_code/goroutine/gorouteUtils"
	"runtime"
	"strconv"
	"time"
)

func goroutineTest() {
	for i := 0; i < 10; i++ {
		fmt.Println("test function ", strconv.Itoa(i))
		time.Sleep(time.Second * 2)
	}
}
func firstGorouteTest() {
	fmt.Println("CPU个数：", runtime.NumCPU())

	// 设置使用多少CPU来运行go程序
	runtime.GOMAXPROCS(runtime.NumCPU())

	go goroutineTest()

	for i := 0; i < 10; i++ {
		fmt.Println("main function ", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	gorouteUtils.GorouteTest1()
}
