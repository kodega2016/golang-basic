package main

import "fmt"

func main() {
	userDoc := map[string]string{
		"name":    "John",
		"address": "Kathmandu",
		"score":   "1200",
	}
	fmt.Println(userDoc)

	// delete(userDoc, "name")
	// updating the map value
	userDoc["name"] = "Khadga Bahadur Shrestha"

	name, isOk := userDoc["name"]
	if isOk {
		fmt.Println("name:", name)
	}

	for key, value := range userDoc {
		fmt.Println(key, value)
	}
}
