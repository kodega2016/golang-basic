package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("You are on macos")
	case "linux":
		fmt.Println("You are on linux")
	default:
		fmt.Println("you are on windows")
	}
}
