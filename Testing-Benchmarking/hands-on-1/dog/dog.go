package dog

// Years converts human years to dog years.
func Years(n int) int {
	return n * 7
}

// YearsTwo converts human years to dog years using a loop.
func YearsTwo(n int) int {
	count := 0
	for range n {
		count += 7
	}
	return count
}
