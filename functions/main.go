package main

import "fmt"

func main() {
	sayHello()
}

func sayHello() {
	defer clear()
	fmt.Println("Hi...")
	logMessage("learning golang for software development")
	total := getSum(10, 20)
	fmt.Println("total:", total)
	fmt.Println(swap("Hello", "World"))
	fmt.Println(getResult(20, 30))
}

func logMessage(message string) {
	fmt.Println("log:", message)
}

func getSum(a, b int) int {
	return a + b
}

func swap(first, second string) (string, string) {
	return second, first
}

func getResult(a, b int) (sum int, mul int) {
	sum = a + b
	mul = a * b
	return
}

func clear() {
	fmt.Println("clearing resources...")
}

func init() {
	fmt.Println("initializing resources...")
}
