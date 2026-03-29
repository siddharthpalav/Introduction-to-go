package main

import (
	"fmt"

	"gihub.com/shadowsid/podcast/auth"
	"gihub.com/shadowsid/podcast/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("shadowsid", "secret")

	session := auth.GetSession()

	fmt.Println(session)

	user := user.User{
		Email: "user@email.com",
		Name:  "John Doe",
	}

	fmt.Println(user)

	color.Red(user.Name)

}
