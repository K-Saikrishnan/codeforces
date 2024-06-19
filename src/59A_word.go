// https://codeforces.com/problemset/problem/59/A

package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		word     string
		lowerCnt int
		upperCnt int
	)

	fmt.Scan(&word)

	for _, c := range word {
		if c >= 'a' && c <= 'z' {
			lowerCnt++
		} else {
			upperCnt++
		}
	}

	if lowerCnt >= upperCnt {
		fmt.Println(strings.ToLower(word))
	} else {
		fmt.Println(strings.ToUpper(word))
	}
}
