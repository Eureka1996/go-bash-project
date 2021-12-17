package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var count int
	rand.Seed(time.Now().UnixNano())
	for {
		// time.Now().Unix()：秒
		// time.Now().UnixNano():纳秒

		// 随机生成 1 ~ 100的数
		var num int = rand.Intn(100) + 1
		//fmt.Println(num)
		if num == 99 {
			break
		} else {
			count += 1
		}
	}
	fmt.Println(count)
}
