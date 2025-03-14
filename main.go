package main

import (
	"fmt"

	"haruto.dev/go-circular/post"
	"haruto.dev/go-circular/user"
)

func main() {
	fmt.Println("Hello")

	user := user.User{
		Name: "Haruto",
	}

	post := post.Post{
		Title:  "Hello",
		Author: user,
	}

	fmt.Println(user)
	fmt.Println(post)
}
