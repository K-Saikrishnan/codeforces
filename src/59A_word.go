// https://codeforces.com/problemset/problem/59/A

package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		word    string
		lowerCt int
		upperCt int
	)

	fmt.Scan(&word)

	for _, char := range word {
		if char >= 'a' && char <= 'z' {
			lowerCt++
		} else {
			upperCt++
		}
	}

	if lowerCt >= upperCt {
		fmt.Println(strings.ToLower(word))
	} else {
		fmt.Println(strings.ToUpper(word))
	}
}
