// https://codeforces.com/problemset/problem/110/A

package main

import "fmt"

func main() {
	var n, luckyNumCnt int

	fmt.Scan(&n)

	for n > 0 {
		if n%10 == 4 || n%10 == 7 {
			luckyNumCnt++
		}
		n /= 10
	}

	if luckyNumCnt == 4 || luckyNumCnt == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
