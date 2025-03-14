package packageB

import (
	"example.com/go-circular/packageA"
)

type B struct {
	A *packageA.A
}
