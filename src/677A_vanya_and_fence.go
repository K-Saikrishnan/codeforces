// https://codeforces.com/problemset/problem/677/A

package main

import "fmt"

func main() {
	var friendsCnt, maxHeight, height, minWidth int

	fmt.Scan(&friendsCnt, &maxHeight)

	minWidth = friendsCnt

	for range friendsCnt {
		fmt.Scan(&height)

		if height > maxHeight {
			minWidth++
		}
	}

	fmt.Println(minWidth)
}
