// https://codeforces.com/problemset/problem/236/A

package main

import "fmt"

func main() {
	var name string

	fmt.Scan(&name)

	unique := make(map[rune]bool)
	for _, char := range name {
		unique[char] = true
	}

	if len(unique)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
