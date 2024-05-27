// 05/26/2024
package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Please input a word: ")
	fmt.Scan(&input)
	input = strings.ToLower(input)
	if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found!")
	}

}
