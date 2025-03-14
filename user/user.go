package user

import "haruto.dev/go-circular/post"

type User struct {
	Name  string
	Posts []post.Post
}
