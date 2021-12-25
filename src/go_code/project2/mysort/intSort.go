package mysort

import (
	"fmt"
	"math/rand"
	"project2/src/go_code/project2/entities"
	"sort"
	"time"
)

func SortIntTest() {

	intSlice := []int{3, 2, -2, 6, 4}

	sort.Ints(intSlice)
	fmt.Println(intSlice)
}

func SortStructTest() {
	var heros entities.HeroSlice
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		hero := entities.Hero{
			Name: fmt.Sprintf("英雄-%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	fmt.Println(heros)
	sort.Sort(heros)
	fmt.Println(heros)
}
