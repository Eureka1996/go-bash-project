package timer_utils

import (
	"fmt"
	"time"
)

/**
 * @author wufuqiang
 * @date 2022/6/5 21:52
 */
// 阻塞指定时间
func TimerTest1() {
	timer := time.NewTimer(time.Second * 10)
	fmt.Println("当前时间：", time.Now())
	t1 := <-timer.C // 阻塞的，指定时间到
	fmt.Println("t1:", t1)

	// 如果只是单纯的等待，可以使用time.Sleep来实现。
	timer2 := time.NewTimer(time.Second * 5)
	<-timer2.C
	fmt.Println("5s之后")

	time.Sleep(time.Second * 5)
	fmt.Println("再一次5s后")

	<-time.After(time.Second * 5) // time.After函数的返回值是chan Time
	fmt.Println("再再一次5s后")

}

func TimerTest2() {
	timer := time.NewTimer(time.Second * 2)

	go func() {
		<-timer.C
		fmt.Println("Timer expired")
	}()
	// stop后，上述协程直接不执行
	stop := timer.Stop()

	if stop {
		fmt.Println("Timer stopped")
	}

	<-time.After(time.Second * 5)
	fmt.Println("主方法")
}
