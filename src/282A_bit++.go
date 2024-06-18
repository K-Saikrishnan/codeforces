// https://codeforces.com/problemset/problem/282/A

package main

import "fmt"

func main() {
	var (
		n    int
		x    int
		expr string
	)

	fmt.Scan(&n)

	for range n {
		fmt.Scan(&expr)

		if expr == "++X" || expr == "X++" {
			x++
		} else {
			x--
		}
	}

	fmt.Println(x)
}
