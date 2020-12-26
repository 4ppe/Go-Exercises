package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Print daily requests
//
//  You've got request logs of a web server. The log data
//  contains 8-hourly totals per each day. It is stored
//  in the `reqs` slice.
//
//  Find and print the total requests per day, as well as
//  the grand total.
//
//  See the `reqs` slice and the steps in the code below.
//
//
// RESTRICTIONS
//
//   1. You need to produce the daily slice, don't just loop
//      and print the element totals directly. The goal is
//      gaining more experience in slice operations.
//
//   2. Your code should work even if you add to or remove the
//      existing elements from the `reqs` slice.
//
//      For example, after solving the exercise, try it with
//      this new data:
//
var reqs = []int{
	500, 600, 250,
	200, 400, 50,
	900, 800, 600,
	750, 250, 100,
	150, 654, 235,
	320, 534, 765,
	121, 876, 285,
	543, 642,
	// the last element is missing (your code should be able to handle this)
	// that is why you shouldn't calculate the `size` below manually.
}

//
//      The grand total of the new data should be 10525.
//
//
// EXPECTED OUTPUT
//
//   Please run `solution/main.go` to see the expected
//   output.
//
//   go run solution/main.go
//
// ---------------------------------------------------------

func main() {
	// There are 3 request totals per day (8-hourly)
	const N = 3

	// DAILY REQUESTS DATA (8-HOURLY TOTALS PER DAY)
	// reqs := []int{
	// 	500, 600, 250, // 1st day: 1350 requests
	// 	200, 400, 50, // 2nd day: 650 requests
	// 	900, 800, 600, // 3rd day: 2300 requests
	// 	750, 250, 100, // 4th day: 1100 requests
	// 	// grand total: 5400 requests
	// }

	// ================================================
	// #1: Make a new slice with the exact size needed.

	_ = reqs // remove this when you start

	size := len(reqs) / 3 // you need to find the size.
	if len(reqs)%3 != 0 {
		size++
	}

	daily := make([][]int, 0, size)
	// ================================================

	// ================================================
	// #2: Group the `reqs` per day into the slice: `daily`.
	//
	// So the daily will be:
	// [
	//  [500, 600, 250]
	//  [200, 400, 50]
	//  [900, 800, 600]
	//  [750, 250, 100]
	// ]

	for i := 0; i < len(reqs); i += 3 {

		// adds the last day
		if i > len(reqs)-3 {
			daily = append(daily, reqs[i:])
			break
		}

		daily = append(daily, reqs[i:N+i])

	}

	// ================================================
	// #3: Print the results

	// Print a header
	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))

	// Loop over the daily slice and its inner slices to find
	// the daily totals and the grand total.

	var grand int
	var total int

	for day, dReqs := range daily {
		fmt.Println()
		total = 0
		for _, v := range dReqs {
			total += v
			fmt.Printf("%-10d%-10d\n", day+1, v)
		}
		fmt.Printf("%10s%d\n", "TOTAL: ", total)
		grand += total
	}

	fmt.Printf("\n%10s%d\n", "GRAND: ", grand)

	// ================================================
}
