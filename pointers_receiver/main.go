package main

import "fmt"

func main() {
	c := Counter{
		count: 0,
	}

	fmt.Println("Initial count:", c.count)
	c.increment()
	fmt.Println("Count after increment:", c.count)
}

type Counter struct {
	count int
}

func (c *Counter) increment() {
	c.count++
}
