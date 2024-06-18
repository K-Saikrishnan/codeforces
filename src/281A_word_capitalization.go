// https://codeforces.com/problemset/problem/281/A

package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string

	fmt.Scan(&word)

	capitalizedWord := strings.ToUpper(word[:1]) + word[1:]

	fmt.Println(capitalizedWord)
}
