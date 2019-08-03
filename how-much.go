/*
I always thought that my old friend John was rather richer than he looked,
but I never knew exactly how much money he actually had. One day (as I was
plying him with questions) he said: "Imagine I have between m and n Zloty
(or did he say Quetzal? I can't remember!)

If I were to buy 9 cars costing c each, I'd only have 1 Zlotty (or was it Meticals?) left.

And if I were to buy 7 boats at b each, I'd only have 2 Ringglets (or was it Zlotty?) left.

Could you tell me in each possible case:

how much money f he could possibly have
the cost c of a car
the cost b of a boat?
So, I will have a better idea about his fortune. Note that if m-n is big enough,
you might have a lot of possible answers.

Each answer will be given as ["M: f", "B: b", "C: c"] and all the answers
as [ ["M: f", "B: b", "C: c"] ... ]. M stands for "Money", B for boats, C for cars.

m and n are positive or null integers with m <= n or m >= n, m and n inclusive.

##Examples:

howmuch(1, 100) => [["M: 37", "B: 5", "C: 4"], ["M: 100", "B: 14", "C: 11"]]
howmuch(1000, 1100) => [["M: 1045", "B: 149", "C: 116"]]
howmuch(10000, 9950) => [["M: 9991", "B: 1427", "C: 1110"]]
howmuch(0, 200) => [["M: 37", "B: 5", "C: 4"], ["M: 100", "B: 14", "C: 11"], ["M: 163", "B: 23", "C: 18"]]
Explanation of howmuch(1, 100):

In the first answer his possible fortune is 37 so if he buys 7 boats each
worth 5 it remains 37 - 7 * 5 = 2, if he buys 9 cars worth 4 each it remains 37 - 9 * 4 = 1.
The same with f = 100: 100 - 7 * 14 = 2 and 100 - 9 * 11 = 1.

Note
See "Sample Tests" to know the format of the return.
*/

package main

import (
	"fmt"
)

func main() {
	result := HowMuch(605, 846)
	fmt.Println(result)
}

func HowMuch(m int, n int) [][3]string {
	if m == 0 && n == 0 {
		return [][3]string{{"M:0", "B:0", "C:0"}}
	}

	max := max(m, n)
	maxCarPrice := max/9 + 1
	maxBoatPrice := max/7 + 2

	fmt.Println("maxCarPrice = ", maxCarPrice)
	fmt.Println("maxBoatPrice= ", maxBoatPrice)

	min := min(m, n)
	minCarPrice := min/9 + 1
	minBoatPrice := min/7 + 2

	fmt.Println("minCarPrice = ", minCarPrice)
	fmt.Println("minBoatPrice= ", minBoatPrice)

	var solutions [][3]string

	for i := minCarPrice; i <= maxCarPrice; i++ {
		for j := minBoatPrice; j <= maxBoatPrice; j++ {
			if 9*i-7*j == 1 && (9*i+1 <= max && 9*i+1 >= min) {
				var solution [3]string
				solution[0] = fmt.Sprintf("M: %d", 9*i+1)
				solution[1] = fmt.Sprintf("B: %d", j)
				solution[2] = fmt.Sprintf("C: %d", i)
				solutions = append(solutions, solution)
			}
		}
	}
	return solutions
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
