package packageA

import (
	"example.com/go-circular/packageB"
)

type A struct {
	B *packageB.B
}
