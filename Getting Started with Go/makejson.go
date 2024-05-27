package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please input your name:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Please input your address:")
	scanner.Scan()
	address := scanner.Text()

	person := map[string]string{
		"name":    name,
		"address": address,
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Something wrong")
		return
	} else {
		fmt.Printf("Json:", jsonData)
	}
}
