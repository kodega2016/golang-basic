package main

import "fmt"

func main() {
	kode := User{
		"Khadga Bahadur Shrestha",
		25,
	}
	fmt.Println(kode)
}

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", u.Name, u.Age)
}
