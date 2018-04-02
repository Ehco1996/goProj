package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	l1 := [...]int{1, 2, 3}
	l11 := [...]int{1, 2, 3}
	l2 := [...]int{1, 2, 3}
	l3 := [...]int{1, 2, 3, 4}
	reverse(l1[:])

	fmt.Println(equal(l1[:], l11[:]))
	fmt.Println(equal(l1[:], l2[:]))
	fmt.Println(equal(l11[:], l2[:]))
	fmt.Println(equal(l1[:], l3[:]))

}
