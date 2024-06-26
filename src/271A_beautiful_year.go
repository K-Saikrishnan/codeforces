// https://codeforces.com/problemset/problem/271/A

package main

import "fmt"

func main() {
	var year int

	fmt.Scan(&year)

	for {
		year++

		a := year / 1000
		b := year / 100 % 10
		c := year / 10 % 10
		d := year % 10

		if a != b && a != c && a != d && b != c && b != d && c != d {
			break
		}
	}

	fmt.Println(year)
}
