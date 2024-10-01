package piscine

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	if is_prime(nb) {
		return nb
	}
	return FindPrevPrime(nb - 1)
}

func is_prime(n int) bool {
	if n < 2 {
		return false
	}
	i := 2
	for i <= n/i {
		if n % i == 0 {
			return false
		}
		i++
	}
	return true
}
