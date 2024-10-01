package piscine

func HashCode(dec string) (s string) {
	x := 0
	for i := 0; i < len(dec); i++ {
		x = (int(dec[i]) + len(dec)) % 127
		if x >= 0 && x < 127 {
			s += string(x)
		} else {
			s += string(dec[i] + 33)
		}
	}
	return
}
