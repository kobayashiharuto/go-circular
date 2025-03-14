package post

import "haruto.dev/go-circular/user"

type Post struct {
	Title  string
	Author user.User
}
