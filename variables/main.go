package main

import "fmt"

func main() {
	var name string
	name = "Software development with golang"
	fmt.Println(name)

	course := "Docker for devops engineer"
	fmt.Println(course)

	var (
		username string = "kodega2016"
		role     string = "Software Developer"
		age      int    = 27
	)
	fmt.Println(username, role, age)

	project_name, project_owner := "Kart", "Jack"
	fmt.Println(project_name, project_owner)
}
