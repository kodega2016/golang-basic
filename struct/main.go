package main

import "fmt"

func main() {
	user := User{
		Name: "Khadga Shrestha",
		Age:  28,
	}
	ptr := &user
	ptr.Age = 29
	fmt.Println(user.Name, user.Age)
}

type User struct {
	Name string
	Age  int
}
