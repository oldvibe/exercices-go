package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]

	res := ""
	inWord := false

for i, ch := range arg {
	if ch != ' ' && ch != '\t' {
		if inWord && i- 1 > 0 && (arg[i - 1] == ' ' || arg[i - 1] == '\t') {
			res += "   "
		} 
		res += string(ch)
		inWord = true
		 
	}
}
	fmt.Println(res)
}
