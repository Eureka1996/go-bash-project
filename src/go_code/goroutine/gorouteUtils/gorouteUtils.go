package gorouteUtils

import (
	"fmt"
	"time"
)

func WriteData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("writeData = ", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func ReadData(intChan chan int, exitChan chan bool) {
	for {
		value, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readData 读取到的数据：", value)
		time.Sleep(time.Second * 5)
	}
	exitChan <- true
	close(intChan)
}

func GorouteTest1() {

	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go WriteData(intChan)
	go ReadData(intChan, exitChan)

	for {
		canExit := <-exitChan
		fmt.Println("主线程去判断是否能结束")
		if canExit {
			break
		}
	}

}
