package gorouteUtils

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @author wufuqiang
 * @date 2022/6/5 14:18
 */

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
	time.Sleep(time.Second * 10)
}

func WaitGroupTest1() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}

	wg.Wait() //等待所有登记的goroutine都结束
	fmt.Println("全部结束。")
}
