package main

import (
	"fmt"
	"strconv"
)

type numbers []int

func main() {
	input := numbers{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	input.print()

}

func (n numbers) print() {
	for _, number := range n {
		fmt.Println(strconv.Itoa(number) + " is " + isEvenOdd(number))
	}
}

func isEvenOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}
