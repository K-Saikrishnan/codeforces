// https://codeforces.com/problemset/problem/977/A

package main

import "fmt"

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	for range k {
		if n%10 == 0 {
			n /= 10
		} else {
			n--
		}
	}

	fmt.Println(n)
}
