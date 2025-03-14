package user

import "haruto.dev/go-circular/post"

type User struct {
	ID    string
	Name  string
	Posts []post.Post
}
