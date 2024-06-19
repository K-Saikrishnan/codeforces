// https://codeforces.com/problemset/problem/339/A

package main

import (
	"fmt"
	"strings"
)

func main() {
	var calc string

	fmt.Scan(&calc)

	var one, two, three int

	for _, c := range calc {
		if c == '1' {
			one++
		} else if c == '2' {
			two++
		} else if c == '3' {
			three++
		}
	}

	rearranged := strings.Repeat("1+", one) + strings.Repeat("2+", two) + strings.Repeat("3+", three)

	fmt.Println(strings.Trim(rearranged, "+"))
}
