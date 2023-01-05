package json_utils

import (
	"encoding/json"
	"fmt"
)

/**
 * @author wufuqiang
 * @date 2023/1/5 23:37
 */

func Join() {
	typeToGoType := map[float64]string{
		1: "int32",
		2: "float32",
		3: "string",
	}
	var str string = "[  {    \"Name\": \"os\",    \"Type\": 3  },  {    \"Name\": \"curr_lqi_score\",    \"Type\": 2  },  {    \"Name\": \"avg_lqi_score\",    \"Type\": 2  },  {    \"Name\": \"best_lqi_score\",    \"Type\": 2  },  {    \"Name\": \"curr_lqi_rank\",    \"Type\": 1  },  {    \"Name\": \"curr_pull_stream_succ_rate_score\",    \"Type\": 2  },  {    \"Name\": \"goal_pull_stream_succ_rate_score\",    \"Type\": 2  },  {    \"Name\": \"avg_pull_stream_succ_rate_score\",    \"Type\": 2  },  {    \"Name\": \"best_pull_stream_succ_rate_score\",    \"Type\": 2  },  {    \"Name\": \"curr_first_frame_duration_score\",    \"Type\": 2  },  {    \"Name\": \"goal_first_frame_duration_score\",    \"Type\": 2  },  {    \"Name\": \"avg_first_frame_duration_score\",    \"Type\": 2  },  {    \"Name\": \"best_first_frame_duration_score\",    \"Type\": 2  },  {    \"Name\": \"curr_lagging_score\",    \"Type\": 2  },  {    \"Name\": \"goal_lagging_score\",    \"Type\": 2  },  {    \"Name\": \"avg_lagging_score\",    \"Type\": 2  },  {    \"Name\": \"best_lagging_score\",    \"Type\": 2  },  {    \"Name\": \"curr_latency_score\",    \"Type\": 2  },  {    \"Name\": \"goal_latency_score\",    \"Type\": 2  },  {    \"Name\": \"avg_latency_score\",    \"Type\": 2  },  {    \"Name\": \"best_latency_score\",    \"Type\": 2  },  {    \"Name\": \"curr_visual_quality_score\",    \"Type\": 2  },  {    \"Name\": \"goal_visual_quality_score\",    \"Type\": 2  },  {    \"Name\": \"avg_visual_quality_score\",    \"Type\": 2  },  {    \"Name\": \"best_visual_quality_score\",    \"Type\": 2  },  {    \"Name\": \"curr_pull_stream_fail_rate_no_frame\",    \"Type\": 2  },  {    \"Name\": \"goal_pull_stream_fail_rate_no_frame\",    \"Type\": 2  },  {    \"Name\": \"curr_pull_stream_fail_rate_quick_flip\",    \"Type\": 2  },  {    \"Name\": \"goal_pull_stream_fail_rate_quick_flip\",    \"Type\": 2  },  {    \"Name\": \"curr_pull_stream_fail_rate_no_stream\",    \"Type\": 2  },  {    \"Name\": \"goal_pull_stream_fail_rate_no_stream\",    \"Type\": 2  },  {    \"Name\": \"curr_pull_stream_fail_rate_error\",    \"Type\": 2  },  {    \"Name\": \"goal_pull_stream_fail_rate_error\",    \"Type\": 2  },  {    \"Name\": \"curr_sdk_dns_analysis_duration\",    \"Type\": 2  },  {    \"Name\": \"goal_sdk_dns_analysis_duration\",    \"Type\": 2  },  {    \"Name\": \"curr_player_dns_analysis_duration\",    \"Type\": 2  },  {    \"Name\": \"goal_player_dns_analysis_duration\",    \"Type\": 2  },  {    \"Name\": \"curr_tcp_connect_duration\",    \"Type\": 2  },  {    \"Name\": \"goal_tcp_connect_duration\",    \"Type\": 2  },  {    \"Name\": \"curr_tcp_first_package_duration\",    \"Type\": 2  },  {    \"Name\": \"goal_tcp_first_package_duration\",    \"Type\": 2  },  {    \"Name\": \"curr_first_video_package_duration\",    \"Type\": 2  },  {    \"Name\": \"goal_first_video_package_duration\",    \"Type\": 2  },  {    \"Name\": \"curr_first_video_frame_decode_duration\",    \"Type\": 2  },  {    \"Name\": \"goal_first_video_frame_decode_duration\",    \"Type\": 2  },  {    \"Name\": \"curr_first_frame_render_duration\",    \"Type\": 2  },  {    \"Name\": \"goal_first_frame_render_duration\",    \"Type\": 2  },  {    \"Name\": \"curr_first_frame_duration_pct\",    \"Type\": 3  },  {    \"Name\": \"goal_first_frame_duration_pct\",    \"Type\": 3  },  {    \"Name\": \"curr_sei_delay_pct\",    \"Type\": 3  },  {    \"Name\": \"goal_sei_delay_pct\",    \"Type\": 3  },  {    \"Name\": \"curr_vqscore\",    \"Type\": 2  },  {    \"Name\": \"goal_vqscore\",    \"Type\": 2  },  {    \"Name\": \"curr_brightness\",    \"Type\": 2  },  {    \"Name\": \"goal_brightness\",    \"Type\": 2  },  {    \"Name\": \"curr_overexposure\",    \"Type\": 2  },  {    \"Name\": \"goal_overexposure\",    \"Type\": 2  },  {    \"Name\": \"curr_noise\",    \"Type\": 2  },  {    \"Name\": \"goal_noise\",    \"Type\": 2  },  {    \"Name\": \"curr_colorfulness\",    \"Type\": 2  },  {    \"Name\": \"goal_colorfulness\",    \"Type\": 2  },  {    \"Name\": \"curr_download_speed\",    \"Type\": 2  },  {    \"Name\": \"goal_download_speed\",    \"Type\": 2  },  {    \"Name\": \"curr_vqscore_by_resolution\",    \"Type\": 3  },  {    \"Name\": \"goal_vqscore_by_resolution\",    \"Type\": 3  },  {    \"Name\": \"curr_download_speed_by_resolution\",    \"Type\": 3  },  {    \"Name\": \"goal_download_speed_by_resolution\",    \"Type\": 3  },  {    \"Name\": \"curr_vv_distribution_resolution\",    \"Type\": 3  },  {    \"Name\": \"goal_vv_distribution_resolution\",    \"Type\": 3  }]"
	mapArray := make([]map[string]interface{}, 0)
	json.Unmarshal([]byte(str), &mapArray)
	fmt.Println(mapArray)
	for _, m := range mapArray {

		fmt.Printf("%s %s `gorm:\"column:%s\"`\n", camelString(m["Name"].(string)), typeToGoType[m["Type"].(float64)], m["Name"])

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
