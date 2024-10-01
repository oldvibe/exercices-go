package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		fmt.Println()
		return 
	}

	count := 0
	res := ""
	for i := 0; i < len(arg[0]); i++{
	for	j := count; j < len(arg[1]); j++{
			if arg[0][i] == arg[1][j] {
				res += string(arg[0][i])
				count++
				break
			}
		}
	}
	fmt.Println(res)
	if arg[0] == res  {
		fmt.Println(1)
	} else {
		fmt.Println(0)

	}
	
}
