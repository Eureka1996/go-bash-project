package main

import (
	"fmt"
	"project2/src/go_code/my_array/slice_utils"
)

var (
	arr1   = [3]int{1, 2, 3}
	slice1 = []int{1, 2, 3, 4, 5}
)

func main() {
	//arr(arr1)
	slice4 := []int{1, 2, 3, 4, 5}
	Slicep(slice4)
	//
	//for a := range arr1 {
	//	fmt.Println(a)
	//}
	//fmt.Println("---------------------")

	fmt.Println(slice4)
	for _, v := range slice4 {
		fmt.Println(v)
	}

	slice_utils.Test1()
}

// 数组是值传递
func arr(array [3]int) {
	array[0] = 10
	fmt.Printf("%p\n", &arr1)
	fmt.Printf("%p\n", &array)
}

func Slicep(slice2 []int) {
	slice2[0] = 100
	//fmt.Printf("%p\n", &slice1)
	//fmt.Printf("%p\n", &slice2)
}
