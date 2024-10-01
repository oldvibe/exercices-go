package piscine

import (
	//"piscine"

	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	base := "0123456789abcdef"
	j := 1
	for i := 0; i < len(arr); i++ {
		z01.PrintRune(rune(base[arr[i]/16]))
		z01.PrintRune(rune(base[arr[i]%16]))

		if j%4 != 0 && j != len(arr) {
			z01.PrintRune(' ')
		}
		if j%4 == 0 {
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
