package main

func atoi(s string) int {
	res := 0
	sign := 1

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
			res = res * 10 + (int(s[i] - 48))
		}
	}
	return res * sign
}

