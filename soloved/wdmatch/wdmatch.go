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
	s1 := arg[0]
	s2 := arg[1]
	count := 0
	var res [] byte 
	for i := 0; i < len(s1); i++{
		for	j := count; j < len(s2); j++{
			count++
			if arg[0][i] == s2[j] {
				res = append(res, s2[j])
				break
			}
		}
	}
	//fmt.Println(res)
	if string(res) == arg[0]  {
		fmt.Print(arg[0])
	} 
	fmt.Println()
	
}
