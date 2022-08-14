package main

import (
	"fmt"
	"math/rand"
	"time"
)

var queue = make(chan int32)

func doStuff(n int32, m string, a bool) int32 {
	r := rand.Int31n(13)

	fmt.Printf("Doing stuff: %v [%d]\n", m, r)

	time.Sleep(time.Duration(r) * time.Second)

	ret := n + r

	if a {
		queue <- ret
	}

	return ret
}

func hello() {
	fmt.Println("Hello World")
}

func main() {
	defer hello()

	r1 := doStuff(1, "Hello", false)
	r2 := doStuff(3, "Three", false)
	r3 := doStuff(9, "NINER", false)

	fmt.Printf("Sync Results: %v, %v, %v\n", r1, r2, r3)

	go doStuff(7, "ASYNC 7", true)
	go doStuff(13, "ASYNC 13", true)
	go doStuff(69, "ASYNC 69", true)

	r4 := <-queue
	r5 := <-queue
	r6 := <-queue

	fmt.Printf("Async Results: %v, %v, %v\n", r4, r5, r6)
}
