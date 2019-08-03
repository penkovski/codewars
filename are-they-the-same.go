/*
Given two arrays a and b write a function comp(a, b) (compSame(a, b) in Clojure)
that checks whether the two arrays have the "same" elements, with the same
multiplicities. "Same" means, here, that the elements in b are the elements
in a squared, regardless of the order.

Examples
Valid arrays
a = [121, 144, 19, 161, 19, 144, 19, 11]
b = [121, 14641, 20736, 361, 25921, 361, 20736, 361]
comp(a, b) returns true because in b 121 is the square of 11, 14641 is the
square of 121, 20736 the square of 144, 361 the square of 19, 25921 the square
of 161, and so on. It gets obvious if we write b's elements in terms of squares:

a = [121, 144, 19, 161, 19, 144, 19, 11]
b = [11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19]
Invalid arrays
If we change the first number to something else, comp may not return true anymore:

a = [121, 144, 19, 161, 19, 144, 19, 11]
b = [132, 14641, 20736, 361, 25921, 361, 20736, 361]
comp(a,b) returns false because in b 132 is not the square of any number of a.

a = [121, 144, 19, 161, 19, 144, 19, 11]
b = [121, 14641, 20736, 36100, 25921, 361, 20736, 361]
comp(a,b) returns false because in b 36100 is not the square of any number of a.

Remarks
a or b might be [] (all languages except R, Shell). a or b might be nil or null
or None or nothing (except in Haskell, Elixir, C++, Rust, R, Shell).

If a or b are nil (or null or None), the problem doesn't make sense so return false.

If a or b are empty the result is evident by itself.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	//a := []int{26, -61, 62, 8, -61, 72, -45, 68, 20, -43, -29}
	//b := []int{64, 400, 676, 841, 1849, 2025, 3721, 3721, 3844, 4624, 5184}
	a := []int{}
	b := []int{}
	result := Comp(a, b)
	fmt.Println(result)
}

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil {
		return false
	}

	if len(array1) == 0 && len(array2) == 0 {
		return true
	}

	if len(array1) != len(array2) {
		return false
	}

	for i := 0; i < len(array1); i++ {
		if array1[i] < 0 {
			array1[i] = -array1[i]
		}
	}

	sort.Ints(array1)
	sort.Ints(array2)

	for i, v := range array1 {
		if v*v != array2[i] {
			return false
		}
	}

	return true
}
