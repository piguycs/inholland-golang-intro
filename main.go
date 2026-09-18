package main

import (
	"fmt"
	"runtime"
)

func sum(a, b, c int) int {
	return a + b + c
}

func main() {
	var os string
	switch runtime.GOOS {
	case "windows", "darwin":
		os = "StupidOS"
	case "linux":
		os = "GoatOS"
	default:
		os = "hopefully TempleOS"
	}

	greeter := "World"
	fmt.Printf("Hello %s on %s!\n", greeter, os)

	for i := range 5 {
		fmt.Printf("%d\n", i)
	}
}
