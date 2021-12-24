package mysort

import (
	"fmt"
	"sort"
)

func SortIntTest() {

	intSlice := []int{3, 2, -2, 6, 4}

	sort.Ints(intSlice)
	fmt.Println(intSlice)
}

func SortStructTest() {
	//var heros entities.HeroSlice
}
