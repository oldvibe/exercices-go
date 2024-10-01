// package main

// import (
// 	"fmt"
// 	"os"
// )

// var slice = []int{}

// func main() {
// 	arg := os.Args[1:]
// 	if len(arg) != 1 {
// 		fmt.Println()
// 		return
// 	}
// 	fprime(atoi(arg[0]), 2)
// 	fmt.Println(slice)
// }
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) != 2 {
        return
    }
    nb, err := strconv.Atoi(os.Args[1])
    if err != nil || nb < 2 {
        return
    }
    for i := 2; i <= nb; i++ {
        if nb%i == 0 { // 42
            if nb/i != 1 {
                fmt.Printf("%d*", i) //42 / 2 = 21/ 3 = 7 2 * 3 * 7  2*3*7
            } else {
                fmt.Printf("%d", i)
            }
            nb /= i
            i--
        }
    }
    fmt.Println()
}
/*
func FindnextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	if is_prime(nb) {
		return nb
	}
	return FindnextPrime(nb + 1)
}

func fprime(n int, div int) { // 30
	fmt.Println("|", div, n/div, "|")
	if n > 1 {
		slice = append(slice, div)
		fprime(n/div, FindnextPrime(div+1))
	}
}7

func is_prime(n int) bool {
	if n < 2 {
		return false
	}
	i := 2
	for i < n {
		if n%i == 0 {
			return false
		}
		i++
	}
	return true
}

func atoi(s string) int {
	sign := 1
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' {
			continue
		}
		if s[i] == '-' || s[i] == '+' {
			if s[i] == '-' {
				sign = sign * -1
			}
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			res = res*10 + (int(s[i] - 48))
		}
	}
	return res * sign
}
*/