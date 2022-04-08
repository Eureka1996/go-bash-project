package reflectUtils

import (
	"fmt"
	"reflect"
)

func ReflectFirstTest(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType =", rType)

	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v,rVal type = %T\n", rVal, rVal)

}

func ReflectGetStructInfo(i interface{}) {

	typ := reflect.TypeOf(i)
	value := reflect.ValueOf(i)

	kd := value.Kind()
	fmt.Println("value.Kind() = ", kd)
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	fieldNum := value.NumField()
	methodNum := value.NumMethod()
	fmt.Println("属性的个数：", fieldNum)
	fmt.Println("方法的个数：", methodNum)

	// 获取所有属性
	for i := 0; i < fieldNum; i++ {
		fmt.Printf("属性 %d ,值为：%v\n", i, value.Field(i))

		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("属性：%d，json标签为：%v \n", i, tagVal)
		}
	}

}
