package main

import (
	"fmt"
)

// Solution 1 - solution by using array
func convert(s string, numRows int) string {

	if len(s) <= 2 {
		return ""
	}

	strResult := ""
	numColumns := 0

	numDiogonalsElems := numRows - 2
	numElementsPerTwoColumns := numDiogonalsElems + numRows

	if len(s) <= numElementsPerTwoColumns {
		numColumns = numDiogonalsElems + 1
	} else {
		numColumns = (len(s) / numElementsPerTwoColumns) * (numDiogonalsElems + 1)
	}

	if len(s)%numElementsPerTwoColumns != 0 {
		numColumns += (numDiogonalsElems + 1)
	}

	fmt.Println("len(s) - ", len(s))
	fmt.Println("numDiogonalsElems - ", numDiogonalsElems)
	fmt.Println("numElementsPerTwoColumns - ", numElementsPerTwoColumns)
	fmt.Println("numColumns - ", numColumns)

	mySlice := make([][]byte, numRows)
	for i := range mySlice {
		mySlice[i] = make([]byte, numColumns)
	}

	numProcessed := 0
	colProc := 0

	for i := 0; i < numColumns; i++ {
		for j := 0; j < numRows; j++ {

			if numProcessed >= len(s) {
				break
			}

			if i%(numDiogonalsElems+1) == 0 {
				mySlice[j][i] = s[numProcessed]
				numProcessed++
			} else {
				if j == numRows-1 || j == 0 {
					continue
				} else {
					intTemp := numRows - j - 1
					fmt.Println("intTemp - ", intTemp)

					mySlice[numRows-j-1][i] = s[numProcessed]
					numProcessed++
					colProc++
					if colProc >= numDiogonalsElems {
						colProc = 0
						break
					}
					i++
				}
			}
		}
		// i++
	}

	fmt.Println("=====================================")
	for _, row := range mySlice {
		for _, val := range row {
			// fmt.Printf("arr[%d][%d] = %c", i, j, val)
			fmt.Printf("%c", val)
			strResult += string(val)
		}
		// fmt.Printf("_")
	}

	fmt.Println("\n=====================================\n")

	for _, row := range mySlice {
		for _, val := range row {
			fmt.Printf("%c", val)
			// if val != 0 {
			// 	fmt.Println("_")
			// }
		}
		fmt.Printf("\n")
	}

	fmt.Println("\n=====================================")

	return strResult
}

func convert1(s string, numRows int) string {
	strRes := ""

	iteration := (numRows - 1) * 2

	fmt.Println("iteration = ", iteration)

	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += iteration {

			strRes += string(s[j])

			if i < numRows-1 && i > 0 {
				tempVal := j + iteration - 2*i
				fmt.Println(tempVal)
				if (j + iteration - 2*i) < len(s) {
					strRes += string(s[j+iteration-2*i])
				}
			}

		}
	}

	return strRes

}

func main() {
	fmt.Println("The longest substring is - ", convert("PAYPALISHIRING", 4))

	fmt.Println("The longest substring is - ", convert1("PAYPALISHIRING", 4))
}
