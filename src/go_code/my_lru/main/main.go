package main

import (
	"fmt"
	"github.com/hashicorp/golang-lru"
)

func main() {
	fmt.Println("wufuqiang")
	cache, err := lru.New(3)
	fmt.Println(cache, err)
	len := cache.Len()
	fmt.Println(len)

	cache.Add("1", 1)
	cache.Add("2", 2)
	cache.Add("3", 3)
	cache.Add("4", 4)

	get, ok := cache.Get("2")
	if ok {
		fmt.Println(get)
	}

	fmt.Println(cache.Len())
	fmt.Println("-----------------")
	for _, k := range cache.Keys() {
		fmt.Println(k)
		fmt.Println(cache.Get(k))
	}
}
