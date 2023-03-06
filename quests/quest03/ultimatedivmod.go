package piscine

func UltimateDivMod(a *int, b *int) {
	old_a := *a
	*a = *a / *b
	*b = old_a % *b
}
