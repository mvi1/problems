package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	l, maxLen := 0, 0

	seen := make(map[byte]int)

	for r := 0; r < len(s); r++ {
		index, ok := seen[s[r]]

		if ok {
			l = max(l, index)
		}

		seen[s[r]] = r + 1
		maxLen = max(maxLen, r-l+1)
	}

	return maxLen
}

func main() {
	fmt.Println("The longest substring is - ", lengthOfLongestSubstring("abcabcbb"))
}
