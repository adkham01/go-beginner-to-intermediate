// Package word provides functions to count words in a string
package word

import "strings"

// no need to write example for this one
// witing a test for this one is a bonus challenge:harder
// UseCount counts the occurrences of each word in a string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count counts the number of words in a string
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
