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
	result := ""
	mp := map[rune]bool{}
	for _, ch := range arg[1] {
		if !mp[ch] {
			mp[ch] = true
		}
	}
	for _, ch := range arg[0] {
		if mp[ch] {
			result += string(ch)
			mp[ch] = false
		}
	}
	fmt.Println(result)
}
