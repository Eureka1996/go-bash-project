package main

import (
	"fmt"
	"go_base/src/go_code/xinzong_every_day"
)

/**
 * @author wufuqiang
 * @date 2022/7/8 11:18
 */

func function(index, value int) int {
	fmt.Println(index)
	return index
}

func main01() {
	defer function(1, function(3, 0))
	defer function(2, function(4, 0))
}

func main() {
	xinzong_every_day.Test1()
}
