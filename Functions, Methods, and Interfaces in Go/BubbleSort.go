// 05272024

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func BubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				swap(slice, j)
			}
		}
	}
}

func main() {
	fmt.Println("Please input a list with max 10 integer with space separate:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strNums := strings.Split(input, " ")
	print(strNums)
	var nums []int
	for _, strNum := range strNums {
		num, err := strconv.Atoi(strNum)
		if err == nil {
			nums = append(nums, num)
		}
		if len(nums) >= 10 {
			break
		}
	}

	BubbleSort(nums)
	fmt.Println("My int list: ")
	for _, num := range nums {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
