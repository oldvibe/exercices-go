package piscine

import (
	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	base := "0123456789abcdef"
	j := 1
	for _, ch := range arr {
		z01.PrintRune(rune(base[ch / 16]))
		z01.PrintRune(rune(base[ch % 16]))

		if j % 4 != 0 && j < len(arr) {
			z01.PrintRune(' ')
		}
		if j % 4 == 0 {
			z01.PrintRune('\n')
		}
		j++
	}
	z01.PrintRune('\n')
	for _, ch := range arr {
		if ch >= 32 && ch <= 126 {
			z01.PrintRune(rune(ch))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
	
}
