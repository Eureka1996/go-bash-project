package string_utils

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/**
 * @author wufuqiang
 * @date 2022/5/19 22:03
 */

// strings.Split()转接口，再转[]interface{}
func StringSplit2Interface() {

	str := "wufuqiang,wfq"
	var j interface{}
	j = strings.Split(str, ",")
	fmt.Println(j)
	// 失败， interface {} is []string, not []interface {}
	//i := j.([]interface{})
	i := j.([]string)

	for index, value := range i {
		s := value
		fmt.Printf("index: %v,value:%v \n", index, s)
	}
}

// 中文字符串长度
func StringLength() {

	fmt.Println(len("你好")) // 6

	fmt.Println(utf8.RuneCountInString("你好"))   // 2
	fmt.Println(utf8.RuneCountInString("你好ab")) // 4
}
