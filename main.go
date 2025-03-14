package main

import (
	"fmt"

	"example.com/go-circular/packageA"
	"example.com/go-circular/packageB"
)

func main() {
	a := &packageA.A{}
	b := &packageB.B{}
	a.B = b
	b.A = a
	fmt.Println(a, b)
}
