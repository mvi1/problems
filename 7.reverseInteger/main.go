package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {

	intRes := 0

	isNeg := false

	if x < 0 {
		x *= -1
		isNeg = true
		fmt.Println(x)
	}

	tmpX := x

	for {

		tmpRes := tmpX % 10

		if tmpX >= 10 {

			tmpX = tmpX / 10
			tmpRes = intRes*10 + tmpRes*10

		} else {
			intRes += tmpX
			break
		}

		intRes = tmpRes

	}

	if intRes > math.MaxInt32 {
		return 0
	}

	if isNeg {
		return -intRes
	}

	return intRes
}

func main() {
	fmt.Println(math.MaxInt32)
	fmt.Println("Reverse will be - ", reverse(1534236469))

	if 9646324351 > math.MaxInt32 {
		fmt.Println("it's bigger")
	}
}
