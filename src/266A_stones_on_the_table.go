// https://codeforces.com/problemset/problem/266/A

package main

import "fmt"

func main() {
	var (
		n        int
		s        string
		removeCt int
	)

	fmt.Scan(&n, &s)

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			removeCt++
		}
	}

	fmt.Println(removeCt)
}
