package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr1, arr2 := nums1, nums2
	if len(arr1) > len(arr2) {
		arr1, arr2 = nums2, nums1
	}
	totalLen := len(arr1) + len(arr2)
	midTotal := totalLen / 2

	l, r := 0, len(arr1)-1

	fmt.Println("totalLen = ", totalLen)
	fmt.Println("midTotal = ", midTotal)

	leftPos1 := 0
	valLeftPos1, valLeftPos2, valRightPos1, valRightPos2 := 0, 0, 0, 0

	for {

		if l+r < 0 {
			leftPos1 = -1
		} else {
			leftPos1 = (l + r) / 2
		}
		leftPos2 := midTotal - leftPos1 - 2

		fmt.Println("leftPos1 = ", leftPos1)
		fmt.Println("leftPos2 = ", leftPos2)

		if leftPos1 >= 0 {
			valLeftPos1 = arr1[leftPos1]
		} else {
			valLeftPos1 = -math.MaxInt32
		}
		if leftPos2 >= 0 {
			valLeftPos2 = arr2[leftPos2]
		} else {
			valLeftPos2 = -math.MaxInt32
		}
		if leftPos1+1 < len(arr1) {
			valRightPos1 = arr1[leftPos1+1]
		} else {
			valRightPos1 = math.MaxInt32
		}
		if leftPos2+1 < len(arr2) {
			valRightPos2 = arr2[leftPos2+1]
		} else {
			valRightPos2 = math.MaxInt32
		}

		fmt.Println("valLeftPos1 = ", valLeftPos1)
		fmt.Println("valLeftPos2 = ", valLeftPos2)
		fmt.Println("valRightPos1 = ", valRightPos1)
		fmt.Println("valRightPos2 = ", valRightPos2)

		// return float64(2.5)

		fmt.Printf("%d <= %d && %d <= %d\n", valLeftPos1, valRightPos2, valLeftPos2, valRightPos1)

		// return float64(2.5)

		if valLeftPos1 <= valRightPos2 && valLeftPos2 <= valRightPos1 {
			if totalLen%2 == 1 {
				fmt.Println(valRightPos1, valRightPos2)
				return float64(min(valRightPos1, valRightPos2))
			} else {
				return float64(max(valLeftPos1, valLeftPos2)+min(valRightPos1, valRightPos2)) / 2
			}
		} else {
			if valLeftPos1 > valRightPos2 {
				r = leftPos1 - 1
			} else {
				l = leftPos1 + 1
			}
		}

	}

}

func main() {
	nums1 := []int{1, 3, 4}
	nums2 := []int{2, 5}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
