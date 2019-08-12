package main

import (
	"fmt"
)

//str addition
func main() {
	num1 := "2905"
	num2 := "505"

	bs1, bs2 := []byte(num1), []byte(num2)

	if len(bs2) > len(bs1) {
		bs1, bs2 = bs2, bs1
	}
	zero := []byte("0")[0]

	sum := 0
	carry := 0

	result := []byte{}
	for i, j := len(bs1)-1, len(bs2)-1; i >= -1 || j >= -1; i, j = i-1, j-1 {
		sum += carry
		if i >= 0 {
			sum += int(bs1[i] - zero)
		}

		if j >= 0 {
			sum += int(bs2[j] - zero)
		}

		carry = sum / 10

		result = append(result, byte(sum%10)+zero)
		sum = 0

	}

	for i, j := 0, len(result)-1; i <= j; i, j = i+1, j-1 {
		tmp := result[j]
		result[j] = result[i]
		result[i] = tmp
	}

	if result[0] == []byte("0")[0] {
		fmt.Println(string(result[1:]))
		return
	}

	fmt.Println(string(result))

}
