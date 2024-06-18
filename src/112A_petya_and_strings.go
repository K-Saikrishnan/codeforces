// https://codeforces.com/problemset/problem/112/A

package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string

	fmt.Scan(&a, &b)

	a, b = strings.ToLower(a), strings.ToLower(b)

	fmt.Println(strings.Compare(a, b))
}
