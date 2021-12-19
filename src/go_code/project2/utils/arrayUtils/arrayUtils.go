package arrayUtils

import "fmt"

func ArrayForeach(arr [3]int) {

	for index, value := range arr {
		fmt.Printf("arr[%d] = %d\n", index, value)
		arr[index] += 1
	}

}

func ArrayForeach2(arr [2][3]int) {
	for _, value1 := range arr {
		for _, value2 := range value1 {
			fmt.Printf("%d_%d\t", value2, &value2)
		}
		fmt.Println("")
	}
}
