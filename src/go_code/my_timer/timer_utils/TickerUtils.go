package timer_utils

import (
	"fmt"
	"time"
)

func TickerTest1() {
	ticker := time.NewTicker(time.Second)
	counter := 1
	// 定时执行
	for _ = range ticker.C {
		fmt.Printf("Ticker %d %v\n", counter, time.Now())
		counter += 1
		if counter >= 5 {
			break
		}
	}

	ticker.Stop()
}

func TickerTest2() {
	intChannel := make(chan int)
	productTicker := time.NewTicker(time.Second)
	go func() {
		for _ = range productTicker.C {
			//intChannel <- rand.Intn(100)
			select {
			case intChannel <- 1:
			case intChannel <- 2:
			case intChannel <- 3:
			}
		}
	}()

	consumerTicker := time.NewTicker(time.Second * 2)

	for _ = range consumerTicker.C {
		n := <-intChannel
		fmt.Println("num:", n)
	}

}
