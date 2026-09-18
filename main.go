package main

import (
	"fmt"
	"runtime"
)

func sum(a, b, c int) int {
	return a + b + c
}

func main() {
	var osname string
	switch runtime.GOOS {
	case "windows", "darwin":
		osname = "StupidOS"
	case "linux":
		osname = "GoatOS"
	default:
		osname = "hopefully TempleOS"
	}

	greeter := "World"
	fmt.Printf("Hello %s on %s!\n", greeter, osname)

	for i := range 5 {
		fmt.Printf("%d\n", i)
	}
}
