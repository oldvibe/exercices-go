package piscine

import "fmt"

//import "fmt"

func Itoa(n int) string {
	s := ""
	sign := false
	if n < 0 {
		n = -n
		sign = true
	}
	
	if n == 0 {
		fmt.Print("0")
	}
	for n > 0 {
			s = string(rune(n % 10) + 48) + s
			n = n / 10
		
	} 

	if sign {
		s = "-" + s
	}
	
		
	return s
}
