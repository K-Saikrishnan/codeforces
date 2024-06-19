// https://codeforces.com/problemset/problem/791/A

package main

import "fmt"

func main() {
	var limakWt, bobWt, years int

	fmt.Scan(&limakWt, &bobWt)

	for 3*limakWt <= 2*bobWt {
		years++
		limakWt *= 3
		bobWt *= 2
	}

	fmt.Println(years + 1)
}
