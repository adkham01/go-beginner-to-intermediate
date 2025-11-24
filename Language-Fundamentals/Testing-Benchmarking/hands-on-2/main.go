package main

import (
	"fmt"
	"hands-on-2/quote"
	"hands-on-2/word"
	"sort"
)

func main() {
	fmt.Println(word.Count(quote.Quote))

	m := word.UseCount(quote.Quote)

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	for _, k := range keys {
		fmt.Printf("%q: %d\n", k, m[k])
	}
}
