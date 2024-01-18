package helper

import "strconv"

func ParseInt(str string) int {
	if i, err := strconv.Atoi(str); err != nil {
		// ... handle error
		return 0
	} else {
		return i
	}
}

func ParseFloat(str string) float64 {
	if f, err := strconv.ParseFloat(str, 64); err != nil {
		// ... handle error
		return 0
	} else {
		return f
	}
}
