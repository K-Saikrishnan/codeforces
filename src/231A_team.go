// https: //codeforces.com/problemset/problem/231/A

package main

import "fmt"

func main() {
	var probCount, solCount, petya, vasya, tonya int

	fmt.Scan(&probCount)

	for range probCount {
		fmt.Scan(&petya, &vasya, &tonya)

		if petya+vasya+tonya >= 2 {
			solCount++
		}
	}

	fmt.Println(solCount)
}
