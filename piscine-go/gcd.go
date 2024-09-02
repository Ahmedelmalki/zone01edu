package piscine

func Gcd(a, b uint) uint {
	var i uint
	if a == b {
		return a
	}
	if a < b {
		for i = a; i > 1; i-- {
			if a%i == 0 && b%i == 0 {
				break
			}
		}
	} else {
		for i = b; i > 1; i-- {
			if a%i == 0 && b%i == 0 {
				break
			}
		}
	}
	return i
}
