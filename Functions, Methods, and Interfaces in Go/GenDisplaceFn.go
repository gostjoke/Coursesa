package main

import (
	"fmt"
	"math"
	"strconv"
)

//

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v*t + s
	}
	return fn
}

func main() {
	fmt.Printf("Input Acceleration: ")
	var a, v, s string
	fmt.Scan(&a)
	ac, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Printf("please input a number")
		return
	}

	fmt.Printf("Input Vo: ")
	fmt.Scan(&v)
	vo, err := strconv.ParseFloat(v, 64)
	if err != nil {
		fmt.Printf("please input a number")
		return
	}

	fmt.Printf("Input So: ")
	fmt.Scan(&s)
	so, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("please input a number")
		return
	}

	fn := GenDisplaceFn(ac, vo, so)
	fmt.Println(fn)

	var t float64
	fmt.Println("Enter time:")
	fmt.Scan(&t)

	move := fn(t)
	fmt.Println("Move: ", move)
}
