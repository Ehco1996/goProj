package main

import "fmt"

func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	l := []string{"", "", "123"}
	fmt.Println(l)
	fmt.Println(nonempty2(l))

}
