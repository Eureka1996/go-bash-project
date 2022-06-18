package regexp_utils

import (
	"fmt"
	"math"
	"regexp"
	"time"
)

//匹配每一个时间
func GetString(reg, str string) {

	regx := regexp.MustCompile(reg)

	allString := regx.FindAllString(str, -1)

	fmt.Println(allString, len(allString))

	t0, _ := time.ParseInLocation("2006-01-02", allString[0], time.Local)
	t1, _ := time.ParseInLocation("2006-01-02", allString[1], time.Local)
	fmt.Println(math.Abs(float64(t0.Unix()-t1.Unix())) / 60 / 60 / 24)

}

// 正则获取全部时间
func GetAllNum() {
	str := ` code: 508. hello world: {  code: 504    6789`
	//str = ""
	reg := `code: (([0-9])([0-9]+))`
	regx := regexp.MustCompile(reg)
	// 只获取到第一组匹配的字符串，是一个数组，第一个位置是整个匹配的字符串，
	//之后是每一个分组，分组从左到右，从大到小
	submatch := regx.FindStringSubmatch(str)
	fmt.Println(submatch)
	if submatch != nil && len(submatch) >= 2 {
		fmt.Println(submatch[1])
	}

	// allSubmatch为一个二维数组，第一维数组是每一个匹配到的地方
	allSubmatch := regx.FindAllStringSubmatch(str, -1)
	fmt.Println(allSubmatch)
}
