package json_utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

/**
 * @author wufuqiang
 * @date 2023/1/5 23:37
 */

// [{"Name":"wfq","Type":1}]转成go struct
func Join() {
	typeToGoType := map[float64]string{
		1: "int32",
		2: "float32",
		3: "string",
		4: "string",
	}
	var str string = "[\n\t\t{\n\t\t\t\"Name\": \"strategy_name_cn\",\n\t\t\t\"Type\": 3\n\t\t},\n\t\t{\n\t\t\t\"Name\": \"strategy_type_cn\",\n\t\t\t\"Type\": 3\n\t\t},\n\t\t{\n\t\t\t\"Name\": \"strategy_effect\",\n\t\t\t\"Type\": 6\n\t\t}\n\t]"
	mapArray := make([]map[string]interface{}, 0)
	json.Unmarshal([]byte(str), &mapArray)
	fmt.Println(mapArray)
	for _, m := range mapArray {

		fmt.Printf("%s %s `gorm:\"column:%s\"`\n", camelString(m["Name"].(string)), typeToGoType[m["Type"].(float64)], m["Name"])

	}
}

func TableColumn2Struct() {
	strs := []string{"id", "filter_condition", "modifier", "vqi_conclusion", "general_conclusion", "pull_success_rate_conclusion", "pull_success_rate_guide", "first_frame_duration_conclusion", "first_frame_duration_guide", "delay_conclusion", "delay_guide", "lagging_conclusion", "lagging_guide", "visual_quality_conclusion", "visual_quality_guide", "tone_quality_conclusion", "tone_quality_guide", "create_time", "modified_time"}
	for _, s := range strs {
		fmt.Printf("%s *string `json:\"%s,omitempty\" gorm:\"column:%s\"`\n", camelString(s), s, s)
	}
}

func TableColumn2Upper() {
	strs := []string{"id", "filter_condition", "modifier", "vqi_conclusion", "general_conclusion", "pull_success_rate_conclusion", "pull_success_rate_guide", "first_frame_duration_conclusion", "first_frame_duration_guide", "delay_conclusion", "delay_guide", "lagging_conclusion", "lagging_guide", "visual_quality_conclusion", "visual_quality_guide", "tone_quality_conclusion", "tone_quality_guide", "create_time", "modified_time"}
	for _, s := range strs {
		fmt.Printf("%s = \"%s\"\n", strings.ToUpper(s), s)
	}
}

/**
蛇形转驼峰
* @description xx_yy to XxYx_y_y_y_y_y_y_y to XxYY
* @date 2020/7/30
*@params要转换的字符串
* @return string
**/
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
