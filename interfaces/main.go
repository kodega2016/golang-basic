package main

import "fmt"

func main() {
	var auth BaseAuth = Auth{
		username: "admin",
		password: "password",
	}

	isLoggedIn := auth.login()
	if isLoggedIn {
		fmt.Println("Login successful")
	}
}

type BaseAuth interface {
	login() bool
}

type Auth struct {
	username string
	password string
}

// login implements BaseAuth.
func (a Auth) login() bool {
	if a.username == "user" && a.password == "password" {
		return true
	}
	return false
}
