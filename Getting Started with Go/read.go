package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func truncateString(str string, max int) string {
	if len(str) > max {
		return str[:max]
	}
	return str
}

func main() {
	var filename string

	fmt.Print("input the filename")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Print("file cannot open")
		return
	}
	// if file not open will still use memory
	defer file.Close()

	var names []Name
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) == 2 {
			fname := truncateString(parts[0], 20)
			lname := truncateString(parts[1], 20)
			names = append(names, Name{fname, lname})
		}
	}
	fmt.Println("Read name and last name")
	for _, name := range names {
		fmt.Printf("Name:", name.Fname, "Last name:", name.Lname)
	}
}
