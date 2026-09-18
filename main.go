package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LogEntry struct {
	IP     string
	Method string
	Path   string
	Status string
}

func main() {
	// logs.log is from https://pastebin.com/raw/fbW8wUZ7
	file, err := os.OpenFile("logs.log", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	logs := []LogEntry{}

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			// naievily gonna assume only error will be EOF
			break
		}

		words := strings.Fields(string(line))

		IP := words[0]
		Method := words[5][1:]
		Path := words[6]
		Status := words[8]

		logs = append(logs, LogEntry{IP, Method, Path, Status})
	}

	fmt.Printf("Logfile has %d log entries\n", len(logs))
	for i, logEntry := range logs {
		fmt.Printf("%d %s\n", i, logEntry)
	}
}
