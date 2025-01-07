package main

import "fmt"

func findLongest(s string, l int, r int) string {
	strLen := ""

	for l >= 0 && r < len(s) && s[l] == s[r] {
		if r-l+1 > len(strLen) {
			strLen = s[l : r+1]
		}
		l--
		r++
	}

	return strLen
}

func longestPalindrome(s string) string {
	longestStr := ""

	for i := 0; i < len(s); i++ {

		longestOdd := findLongest(s, i, i)
		longestEven := findLongest(s, i, i+1)

		if len(longestOdd) > len(longestStr) {
			longestStr = longestOdd
		}
		if len(longestEven) > len(longestStr) {
			longestStr = longestEven
		}

	}

	return longestStr

}

func main() {
	fmt.Printf("The longest substring is - \"%s\"\n", longestPalindrome("abkazak"))
}
