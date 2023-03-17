package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative"
	}

	if n%6 == 0 {
		return "rock and roll"
	}
	if n%2 == 0 {
		return "rock"
	}
	if n%3 == 0 {
		return "roll"
	}

	return "error: non divisible"
}
