package sort_util

import (
	"fmt"
	"sort"
	"strings"
)

/**
 * @author wufuqiang
 * @date 2023/2/7 13:36
 */

func Sort1() {
	str := []string{"1080p+", "540p", "360p", "1080p", "other"}
	sort.SliceStable(str, func(i, j int) bool {
		var first string = str[i]
		var second string = str[j]
		if str[i] == "other" {
			first = "0"
		}
		if str[j] == "other" {
			second = "0"
		}

		istr := fmt.Sprintf("%10s", first)
		jstr := fmt.Sprintf("%10s", second)

		compare := strings.Compare(istr, jstr)
		return compare > 0
	})
	fmt.Printf("%v", str)
}
