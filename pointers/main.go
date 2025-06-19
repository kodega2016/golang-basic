package main

import "fmt"

func main() {
	a := 20
	p := &a
	fmt.Printf("The address of the variable a is %v and value is %v\n", p, *p)
}
