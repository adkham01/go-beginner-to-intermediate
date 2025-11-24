package main

import "fmt"

func main() {
	check := []int{21, 43, 64, 57, 75}
	fmt.Println(sum(check))

}

func sum(ii []int) int {
	total := 0
	for _, v := range ii {
		total += v
	}

	return total
}
