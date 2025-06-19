package main

import "fmt"

func main() {
	kode := User{
		"Khadga Bahadur Shrestha",
		25,
	}
	kode.printUserInfo()
}

type User struct {
	Name string
	Age  int
}

func (user User) printUserInfo() {
	fmt.Printf("My name is %s and i am %d\n", user.Name, user.Age)
}
