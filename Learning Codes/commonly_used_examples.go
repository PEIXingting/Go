package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// random digit
	t1 := time.Now()
	timeStamp1 := t1.Unix()
	fmt.Println(timeStamp1)
	timeStamp2 := t1.UnixNano()
	fmt.Println(timeStamp2)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	num1 := rand.Intn(10) + 3 // [3,12]
	fmt.Println(num1)

	// get ordered key, value from map
	map1 := make(map[int]string)
	map1[1] = "one"
	map1[2] = "two"
	map1[3] = "three"
	map1[4] = "four"
	map1[5] = "five"

	keys := make([]int, 0, len(map1))
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Println(key, map1[key])
	}
}
