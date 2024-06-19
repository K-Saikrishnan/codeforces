// https: //codeforces.com/problemset/problem/546/A

package main

import "fmt"

func main() {
	var bananaCost, currentMoney, bananaCnt, lendMoney int

	fmt.Scan(&bananaCost, &currentMoney, &bananaCnt)

	if cost := bananaCost * bananaCnt * (bananaCnt + 1) / 2; currentMoney < cost {
		lendMoney = cost - currentMoney
	}

	fmt.Println(lendMoney)
}
