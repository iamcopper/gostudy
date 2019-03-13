package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args {
		fmt.Println(index, "\t", arg)
	}

	fmt.Println("---------------------------")

	for index, arg := range os.Args {
		fmt.Printf("%d\t%s\n", index, arg)
	}
}
