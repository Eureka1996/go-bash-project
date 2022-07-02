package atomic_utils

import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int32 = 100

func add() {
	atomic.AddInt32(&i, 1)
}

func sub() {
	atomic.AddInt32(&i, -1)
}

func AtomicTest1() {
	for i := 0; i < 100; i++ {
		add()
		sub()
	}
	fmt.Println("num:", i)
}

func CompareAndSwapTest() {
	var i int32 = 100
	go func() { // 此处修改后，后面的CAS会失败
		i = 599
	}()
	time.Sleep(time.Second)
	swapInt32 := atomic.CompareAndSwapInt32(&i, 100, 200)
	fmt.Printf("CompareAndSwap返回的结果：%v \n", swapInt32)
	fmt.Printf("i = %v \n", i)
}
