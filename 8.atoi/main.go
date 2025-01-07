package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	intResult := 0
	isNegative := 1

	numMap := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	for i := 0; i < len(s); i++ {

		if s[i] == ' ' && (i == 0 || s[i-1] == ' ') {
			continue
		}

		fmt.Printf("symbol is - %c\n", s[i])

		if s[i] == '-' && (i == 0 || s[i-1] == ' ') {
			isNegative = -1
			continue
		}

		if s[i] == '+' && (i == 0 || s[i-1] == ' ') {
			continue
		}

		tempVal, ok := numMap[s[i]]
		if !ok {
			break
		}

		tempRes := intResult

		intResult = tempRes*10 + tempVal

		if intResult > math.MaxInt32 {
			if isNegative == 1 {
				return math.MaxInt32
			} else {
				return (math.MaxInt32 + 1) * isNegative
			}
		}

	}

	return isNegative * intResult
}

func main() {
	fmt.Println("Number from string - ", myAtoi("-9223372036854775808"))
}
