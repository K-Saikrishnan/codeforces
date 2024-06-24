// https://codeforces.com/problemset/problem/116/A

package main

import "fmt"

func main() {
	var numStops, exitCnt, entryCnt, passengers, maxPassengers int

	fmt.Scan(&numStops)

	for range numStops {
		fmt.Scan(&exitCnt, &entryCnt)

		passengers += entryCnt - exitCnt

		if passengers > maxPassengers {
			maxPassengers = passengers
		}
	}

	fmt.Println(maxPassengers)
}
