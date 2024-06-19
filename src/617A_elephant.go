// https://codeforces.com/problemset/problem/617/A

package main

import "fmt"

func main() {
	var x, steps int

	fmt.Scan(&x)

	if x%5 == 0 {
		steps = x / 5
	} else {
		steps = x/5 + 1
	}

	fmt.Println(steps)
}
