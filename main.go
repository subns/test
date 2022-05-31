package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("vim-go")
	fmt.Println("yeah")
	fmt.Println("foo")

	var userPtr string
	flag.StringVar(&userPtr, "user", "", "user to greet")
	flag.Parse()

	fmt.Println("Current Time: ", currentTime.String())
}
