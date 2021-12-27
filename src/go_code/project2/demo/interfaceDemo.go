package demo

import (
	"fmt"
	"project2/src/go_code/project2/entities"
)

// 参数类型的判断
func TypeJudge(items ...interface{}) {

	for i, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("param #%d is a bool,value is %v\n", i, v)
		case float32, float64:
			fmt.Printf("param #%d is a float,value is %v\n", i, v)
		case int, int64:
			fmt.Printf("param #%d is a int,value is %v\n", i, v)
		case nil:
			fmt.Printf("param #%d is nil,value is %v\n", i, v)
		case string:
			fmt.Printf("param #%d is a string,value is %v\n", i, v)
		case entities.Phone:
			fmt.Printf("param #%d is a Phone,value is %v\n", i, v)
		case *entities.Phone:
			fmt.Printf("param #%d is a *Phone,value is %v\n", i, v)
		default:
			fmt.Printf("param #%d's type is unknow,value is %v\n", i, v)
		}
	}
}
