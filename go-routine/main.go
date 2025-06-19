package main

import (
	"fmt"
	"time"
)

func main() {
	go play()

	play()
}

func play() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("playing...")
	}
}
