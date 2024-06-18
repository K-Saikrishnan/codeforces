// https://codeforces.com/problemset/problem/71/A

package main

import "fmt"

func main() {
	var (
		numWords int
		word     string
	)

	fmt.Scan(&numWords)

	for range numWords {
		fmt.Scan(&word)

		shortWord := word

		if length := len(word); length > 10 {
			shortWord = string(word[0]) + fmt.Sprint(length-2) + string(word[length-1])
		}

		fmt.Println(shortWord)
	}
}
