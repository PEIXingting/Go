package main

import (
	"fmt"
	"math/rand"
	"time"
)

//random digit
func main() {
	t1 := time.Now()
	timeStamp1 := t1.Unix()
	fmt.Println(timeStamp1)
	timeStamp2 := t1.UnixNano()
	fmt.Println(timeStamp2)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	num1 := rand.Intn(10) + 3 // [3,12]
	fmt.Println(num1)
}