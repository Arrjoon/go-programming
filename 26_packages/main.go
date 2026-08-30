package main

import (
	"fmt"

	"github.com/arrjoon/myapplication/auth"
	"github.com/arrjoon/myapplication/user"
)

func main() {
	auth.LoginWithCredentials("arjun", "password")
	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Email: "nepaliarjun@gmail.com",
		Name:  "Arjun",
	}
	fmt.Println(user.Email)
}
