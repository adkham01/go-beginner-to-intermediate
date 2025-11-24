package mystr

import "strings"

func Cat(xs []string) string {
	var result string
	for _, v := range xs {
		result += v
		result += " "
	}
	return result
}

func Join(xs []string) string {
	return strings.Join(xs, " ")
}
