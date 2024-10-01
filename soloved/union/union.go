package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		return
	}
	mp := map[rune]bool{}
	res := arg[0] + arg[1]
	result := ""

	for _, ch := range res {
		if !mp[ch] {
			result += string(ch)
			mp[ch] = true
		}
	}
	fmt.Println(result)
}
