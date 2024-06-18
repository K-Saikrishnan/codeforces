// https://codeforces.com/problemset/problem/50/A

package main

import "fmt"

func main() {
	var m, n int

	fmt.Scan(&m, &n)

	maxDominos := (m * n) / 2

	fmt.Println(maxDominos)
}
