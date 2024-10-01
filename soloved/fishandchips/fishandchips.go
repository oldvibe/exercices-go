package piscine

import (
	"github.com/01-edu/z01"
)


func FishAndChips(n int) string {
	if n < 0 {
		z01.PrintRune('0')
	}
	if n % 2 == 0 && n % 3 == 0 {
		z01.PrintRune('f')
		z01.PrintRune('c')
	} else if n % 3 == 0 {
		z01.PrintRune('c')
	} else if n % 2 == 0 {
		z01.PrintRune('f')

	} else {
		z01.PrintRune('n')

	}
	return ""
}