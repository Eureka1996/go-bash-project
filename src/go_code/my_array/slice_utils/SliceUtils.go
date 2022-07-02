package slice_utils

import "fmt"

var (
	arr1   = [3]int{1, 2, 3}
	slice1 = []int{1, 2, 3, 4, 5}
)

//
func Test1() {
	s1 := []int{1, 2, 3, 4}
	SliceTransaction(s1)
	fmt.Println(s1)
}

func SliceTransaction(s []int) {
	s[0] = 100
}
