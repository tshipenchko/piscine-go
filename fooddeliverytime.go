package piscine

type food struct {
	preptime int
}

var menu = map[string]food{
	"burger":  {15},
	"chips":   {10},
	"nuggets": {12},
}

func FoodDeliveryTime(order string) int {
	value, ok := menu[order]

	if ok {
		return value.preptime
	}
	return 404
}
