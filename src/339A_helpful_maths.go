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

	for i := range calc {
		if calc[i] == '1' {
			one++
		} else if calc[i] == '2' {
			two++
		} else if calc[i] == '3' {
			three++
		}
	}

	rearranged := strings.Repeat("1+", one) + strings.Repeat("2+", two) + strings.Repeat("3+", three)

	fmt.Println(strings.Trim(rearranged, "+"))
}
