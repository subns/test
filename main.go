package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println("yeah")
	fmt.Println("foo")

	var userPtr string
	flag.StringVar(&userPtr, "user", "", "user to greet")
	flag.Parse()
}
