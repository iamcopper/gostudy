package main

import (
	"./op"
	"fmt"
)

func main() {
	a := op.NewInteger()

	a.Add(1)
	fmt.Printf("a.Add(1), a.Value()=%v\n", a.Value())

	// Build failed, can't refer to unexport field and method
	// a.sub(1)
	//fmt.Printf("a.Add(1), a = %v\n", a.Value())
}
