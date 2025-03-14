package main

import (
	"fmt"

	"haruto.dev/go-circular/model"
)

func main() {
	fmt.Println("Hello")

	user := model.User{
		Name:  "Haruto",
		Posts: []model.Post{},
	}

	post := model.Post{
		Title:  "Hello",
		Author: user,
	}

	user.Posts = append(user.Posts, post)

	fmt.Println(user)
	fmt.Println(post)
}
