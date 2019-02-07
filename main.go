package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hi!, you just executed the most useless app ever!")

	path := os.Args[0]
	err := os.Remove(path)

	if err != nil {
		fmt.Println(err)
	}
}
