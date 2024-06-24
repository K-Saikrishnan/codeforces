// https://codeforces.com/problemset/problem/734/A

package main

import "fmt"

func main() {
	var n, antonWins, danikWins int
	var matches, winner string

	fmt.Scan(&n, &matches)

	for _, match := range matches {
		if match == 'A' {
			antonWins++
		} else {
			danikWins++
		}
	}

	if antonWins > danikWins {
		winner = "Anton"
	} else if danikWins > antonWins {
		winner = "Danik"
	} else {
		winner = "Friendship"
	}

	fmt.Println(winner)
}
