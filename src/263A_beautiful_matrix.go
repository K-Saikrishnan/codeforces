// https://codeforces.com/problemset/problem/263/A

package main

import "fmt"

func main() {
	var matrix [5][5]int

	for i := range 5 {
		for j := range 5 {
			fmt.Scan(&matrix[i][j])

			if matrix[i][j] == 1 {
				fmt.Println(abs(i-2) + abs(j-2))
				return
			}
		}
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
