package string_utils

import (
	"fmt"
	"strings"
)

/**
 * @author wufuqiang
 * @date 2022/5/19 22:03
 */

func StringSplit2Interface() {

	str := "wufuqiang,wfq"
	var j interface{}
	j = strings.Split(str, ",")
	fmt.Println(j)
	i := j.([]interface{})

	for index, value := range i {
		s := value
		fmt.Printf("index: %v,value:%v \n", index, s)
	}
}
