package main

import (
	"errors"
	"fmt"
	"runtime"
)

func sum(a, b, c int) int {
	return a + b + c
}

func dv(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Could not divide by zero")
	}
	return a / b, nil
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

	if value, err := dv(1, 0); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Div result: %d\n", value)
	}

	if value, err := dv(4, 2); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Div result: %d\n", value)
	}
}
