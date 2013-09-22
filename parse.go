package calculator

import "strconv"

func Parse(input string) (int64, error) {
	return strconv.ParseInt(input, 10, 64)
}
