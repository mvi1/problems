package main

import "fmt"

func initMatrix(s string, p string) [][]bool {
	f := make([][]bool, len(s)+1)

	for i := range f {
		f[i] = make([]bool, len(p)+1)
	}

	f[0][0] = true

	return f
}

func printMatrix(f [][]bool) {
	for i := range f {
		for j := range f[i] {
			if f[i][j] {
				fmt.Printf("T | ")
			} else {
				fmt.Printf("F | ")
			}

		}
		fmt.Println(" ")
	}
}

func isMatch(s string, p string) bool {
	lens := len(s)
	lenp := len(p)

	f := initMatrix(s, p)

	for i := 0; i < lens+1; i++ {
		for j := 1; j < lenp+1; j++ {

			if i > 0 && (p[j-1] == '.' || s[i-1] == p[j-1]) {
				f[i][j] = f[i-1][j-1]
			} else if p[j-1] == '*' {
				f[i][j] = f[i][j-2]
				if i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') {
					f[i][j] = f[i][j] || f[i-1][j]
					fmt.Println(i, j)

				}
			}

		}
	}

	printMatrix(f)

	return f[lens][lenp]
}

func main() {
	fmt.Println("Is regex match - ", isMatch("aa", "a*"))
	fmt.Println("Is regex match - ", isMatch("aabbcc", ".*"))
	fmt.Println("Is regex match - ", isMatch("aaaaabbbbbbccccc", "a*b*c*"))
}
