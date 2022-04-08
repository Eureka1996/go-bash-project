package main

import (
	"go_base/src/go_code/reflect/reflectUtils"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	student := Student{
		Name: "wufuqiang",
		Age:  22,
	}
	reflectUtils.ReflectFirstTest(student)
	reflectUtils.ReflectGetStructInfo(student)
}
