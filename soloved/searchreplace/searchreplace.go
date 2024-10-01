package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	arr2 := []rune(arg[2])
	for _, ch := range arg[0] {
		for _, ch1 := range arg[1] {
			if ch == ch1 {
				ch = rune(arr2[0])
				z01.PrintRune(ch)
				break
			}
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}
