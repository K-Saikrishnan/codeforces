// https: //codeforces.com/problemset/problem/546/A

package main

import "fmt"

func main() {
	var bananaCost, currentMoney, bananaCt, lendMoney int

	fmt.Scan(&bananaCost, &currentMoney, &bananaCt)

	if cost := bananaCost * bananaCt * (bananaCt + 1) / 2; currentMoney < cost {
		lendMoney = cost - currentMoney
	}

	fmt.Println(lendMoney)
}
