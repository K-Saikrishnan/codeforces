// https: //codeforces.com/problemset/problem/231/A

package main

import "fmt"

func main() {
	var probCt, solCt, petya, vasya, tonya int

	fmt.Scan(&probCt)

	for range probCt {
		fmt.Scan(&petya, &vasya, &tonya)

		if petya+vasya+tonya >= 2 {
			solCt++
		}
	}

	fmt.Println(solCt)
}
