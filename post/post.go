package post

import "haruto.dev/go-circular/common"

type Post struct {
	Title  string
	Author common.UserRef
}
