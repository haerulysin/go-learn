package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}
