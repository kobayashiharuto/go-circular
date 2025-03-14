package main

import (
	"fmt"

	"haruto.dev/go-circular/post"
	"haruto.dev/go-circular/user"
)

func main() {
	fmt.Println("Hello")

	user := &user.User{
		Name:  "Haruto",
		Posts: []post.Post{},
	}

	post := post.Post{
		Title:  "Hello",
		Author: user,
	}

	user.Posts = append(user.Posts, post)

	fmt.Println(user)
	fmt.Println(post)
}
