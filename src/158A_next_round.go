// https://codeforces.com/problemset/problem/158/A

package main

import "fmt"

func main() {
	var n, k, qualifiedCt int

	fmt.Scan(&n, &k)

	scores := make([]int, n)

	for i := range n {
		fmt.Scan(&scores[i])
	}

	for _, score := range scores {
		if score > 0 && score >= scores[k-1] {
			qualifiedCt++
		}
	}

	fmt.Println(qualifiedCt)
}
