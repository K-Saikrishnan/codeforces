// https://codeforces.com/problemset/problem/4/A

package main

import "fmt"

func main() {
	var weight int

	fmt.Scan(&weight)

	if weight%2 == 0 && weight > 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
