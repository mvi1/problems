package main

import "fmt"

func maxArea(height []int) int {
	l := len(height)

	s := 0
	left := 0
	right := l - 1

	for left < right {

		w := right - left
		h := min(height[left], height[right])

		temp := h * w
		s = max(s, temp)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return s
}

func main() {
	fmt.Println("max is - ", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println("max is - ", maxArea([]int{8, 7, 2, 1}))
}
