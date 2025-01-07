package main

import "fmt"

func Pow(x int, y int) int {

	res := 1

	for i := 1; i < y; i++ {
		res *= x
	}

	return res
}

func isPalindrome(x int) bool {
	isItTrue := true
	numDigits := 0

	tmpVal := x
	for {
		tmpVal = tmpVal / 10

		if tmpVal > 0 {
			numDigits++
		} else {
			numDigits++
			break
		}
	}

	fmt.Println(numDigits)
	tenPowered := Pow(10, numDigits)
	fmt.Println(" 10 ^ 6 = ", tenPowered)

	workVal := x

	for i := 0; i < numDigits/2; i++ {

		startVal := (x / Pow(10, numDigits-i)) % 10

		endVal := workVal % 10
		workVal /= 10

		if startVal == endVal {
			continue
		} else {
			isItTrue = false
			break
		}
	}

	return isItTrue
}

func main() {
	fmt.Println("is 123321 Palindrome - ", isPalindrome(123321))
	fmt.Println("is 12321 Palindrome - ", isPalindrome(12321))
	fmt.Println("is 456789 Palindrome - ", isPalindrome(456789))
}
