package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	printMemUsageFlag := flag.Bool("printMemUsage", false, "print the heap memory usage")
	flag.Parse()
	// if printMemUsageFlag is true, print the heap memory usage and exit the program
	if *printMemUsageFlag {
		printMemUsage()
		os.Exit(0)
	}

	currentTime := time.Now()
	fmt.Println("vim-go")
	fmt.Println("yeah")
	fmt.Println("foo")

	var userPtr string
	flag.StringVar(&userPtr, "user", "", "user to greet")
	flag.Parse()
	if userPtr != "" {
		fmt.Printf("Hello, %v!\n", userPtr)
	} else {
		fmt.Println("Hello, Anonymous!")
	}
	fmt.Println("Current Time: ", currentTime.String())
}

// this function prints the heap memory usage in bytes
func printMemUsage() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Heap in use: %d bytes\n", memStats.HeapAlloc)
}
