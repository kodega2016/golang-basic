package main

import "fmt"

func main() {
	users := []string{}
	users = append(users, "sakar", "himal")
	fmt.Println(users)

	skills := make([]string, 3)
	skills = append(skills, "docker", "aws")
	fmt.Println("skills:", skills)
}
