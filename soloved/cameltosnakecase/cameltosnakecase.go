package piscine

//"github.com/01-edu/z01"

func CamelToSnakeCase(s string) string {
	for i := 0; i < len(s); i++ {
		if Maj(s) || Min(s) {
			return s
		}

		for i := 1; i < len(s); i++ {
			if (s[i] >= 'A' && s[i] <= 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z') {
				return s
			} else if  (s[i] >= 'A' && s[i] <= 'Z') && (s[i-1] >= 'a' && s[i-1] <= 'z') {
				s = s[:i] + "_" + s[i:]
			}
		}
	}
	return s
}

func Maj(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			return false
		}
	}
	return true
}

func Min(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return false
		}
	}
	return true
}
