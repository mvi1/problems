package main

import "fmt"

// ############ Solution 1 ############
// func twoSum(nums []int, target int) []int {
// 	var res []int
// 	var res1 []int

// 	hmap := make(map[int]int)

// 	for i := 0; i < len(nums); i++ {
// 		hmap[nums[i]] = i
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		if (target-nums[i]) > 0 && (hmap[target-nums[i]] > 0) {

// 			// fmt.Println(hmap[target-nums[i]])

// 			res1 = append(res1, i)
// 			res1 = append(res1, hmap[target-nums[i]])

// 			break
// 		}
// 	}

// 	fmt.Println(res1)

// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if (nums[i] + nums[j]) == target {
// 				res = append(res, i)
// 				res = append(res, j)
// 			}
// 		}
// 	}

// 	return res
// }

// ############ Final solution ############
func twoSum(nums []int, target int) []int {
	var res []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				res = append(res, i)
				res = append(res, j)
			}
		}
	}

	return res
}

func main() {
	var nums = []int{2, 7, 11, 15, 4}
	var target = 15

	fmt.Println("--->", twoSum(nums, target))
}
