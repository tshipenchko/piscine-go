package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	menu := map[string]food{
		"burger":  {15},
		"chips":   {10},
		"nuggets": {12},
	}

	value, ok := menu[order]

	if ok {
		return value.preptime
	}
	return 404
}
