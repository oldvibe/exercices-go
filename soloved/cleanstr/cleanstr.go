package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main(){
	arg := os.Args[1:]
	if (len(arg) != 1) {
		return 
	}
	for _, ch := range arg[0] {
		if ch < 33 {
			continue
		} else {
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}