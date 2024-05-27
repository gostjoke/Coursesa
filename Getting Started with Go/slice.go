package main

// 05262024
import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)

	for {
		fmt.Printf("Please input an integer or 'X' to leave: ")
		var input string
		_, err := fmt.Scan(&input)
		if input == "X" {
			break
		}
		if err != nil {
			fmt.Println("Something wrong, input again")
			continue
		}
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("not integer, input again")
			continue
		}

		slice = append(slice, num)
		sort.Ints(slice)

		fmt.Println("current slice: ", slice)
	}
	fmt.Println("Program close")
}
