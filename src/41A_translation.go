// https://codeforces.com/problemset/problem/41/A

package main

import "fmt"

func main() {
	var original, translation string

	fmt.Scan(&original, &translation)

	originalLen := len(original)
	translationLen := len(translation)

	if originalLen != translationLen {
		fmt.Println("NO")
		return
	}

	for i := range original {
		if original[i] != translation[translationLen-1-i] {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
