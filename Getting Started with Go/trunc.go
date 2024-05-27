// 05/26/2024
package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64
	fmt.Print("Input a Float number: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Print("wrong: ")
		return
	} else {
		trun_num := math.Trunc(input)
		fmt.Println("", trun_num)
	}

}
