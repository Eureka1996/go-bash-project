package main

import (
	"encoding/json"
	"fmt"
	"go_base/src/go_code/my_json/json_utils"
	"time"
)

/**
 * @author wufuqiang
 * @date 2022/4/27 15:59
 */

func main() {
	json_utils.TableColumn2Upper()
}

func main1() {
	type FruitBasket struct {
		Name    string
		Fruit   []string
		Id      int64 `json:"ref"` // 声明对应的json key
		Created time.Time
	}

	jsonData := []byte(`
    {
        "Name": "Standard",
        "Fruit": [
             "Apple",
            "Banana",
            "Orange"
        ],
        "ref": 999,
        "Created": "2018-04-09T23:00:00Z"
    }`)

	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(basket.Name, basket.Fruit, basket.Id)
	fmt.Println(basket.Created)

	//Test1()
	text2()
}

func Test1() {
	context := make(map[string]interface{})
	context["brokerService"] = "wufuqiang"
	context["timeout"] = 5 * 60 * 1000

	contextMap := make(map[string]interface{})
	contextMap["timeout"] = 300000
	context["context"] = contextMap

	marshal, _ := json.Marshal(context)
	fmt.Println("context:", string(marshal))

}

func text2() {
	template := `{
   "queryType": "groupBy",
    "dataSource": "${table_name}",
    "dimensions": ["${dimension}"],
    "granularity": "${aggregationTime}",
  "aggregations": [
    {
          "fieldName": "count",
          "hidden": true,
          "name": "playing_total_num_groupby",
          "type": "longSum"
        },
        {
      "fieldName": "count",
      "name": "total",
      "type": "longSum"
    }
  ],
  "postAggregations": [],
  "intervals": [
    "${startTime}/${endTime}"
  ],
"context":{
"name":"wufuqi______________ang",
"age":18,
"isMan":true},
  "filter": {
    "type": "and",
    "fields": [
      {
        "dimension": "event_key",
        "type": "selector",
        "value": "play_stop"
      },
      {
        "dimension": "product_line",
        "type": "selector",
        "value": "live"
      },
      {
              "dimension": "is_preview",
              "type": "selector",
              "value": "0"
            },
{
              "dimension": "has_risk",
              "type": "selector",
              "value": "0"
            },
      {
        "dimension": "useful",
        "type": "selector",
        "value": "1"
      }
    ]
  }
}
`
	addContextTimeout(template)
	//fmt.Println("返回值：", str)
}

func addContextTimeout(template string) string {
	templateMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(template), &templateMap)
	if err != nil {
		return template
	} else {
		context, ok := templateMap["context"]
		if ok { // 查询本身包含context
			oldContextMap := context.(map[string]interface{})
			if _, ok := oldContextMap["timeout"]; !ok {
				oldContextMap["timeout"] = 30000
			}
			templateMap["context"] = oldContextMap
		} else { // 查询本身不包含context
			contextMap := make(map[string]interface{})
			contextMap["timeout"] = 300000
			templateMap["context"] = contextMap
		}
	}
	marshal, _ := json.Marshal(templateMap)
	return string(marshal)
}
