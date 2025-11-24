package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	// Open a file
	f, err := os.Open("text.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()

	// The frequency of words in the file
	words, err := freq(f)
	if err != nil {
		log.Fatalf("error from freq in main %s", err)
	}

	// sort the word frequencies

	pairs := sortWordfrequency(words)

	for _, pair := range pairs {
		fmt.Printf("%s \t\t %d\n", pair.Key, pair.Value)
	}

	w, n, err := maxWord(words)
	if err != nil {
		log.Fatalf("error with maxWord %s", err)
	}
	fmt.Printf("%#v has frequency %d", w, n)
}

func freq(r io.Reader) (map[string]int, error) {
	// Create  a map to store word frequencies
	wordFreq := make(map[string]int)

	// Create Scanner to read the file
	s := bufio.NewScanner(r)
	s.Split((bufio.ScanWords))

	for s.Scan() {
		word := strings.ToLower(s.Text())
		wordFreq[word]++
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return wordFreq, nil
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortWordfrequency(m map[string]int) PairList {
	// Convert the map to a pair list
	pairs := make(PairList, len(m))
	i := 0
	for k, v := range m {
		pairs[i] = Pair{k, v}
		i++
	}

	// sort the pair list
	sort.Sort(pairs)
	return pairs
}

func maxWord(m map[string]int) (string, int, error) {
	if len(m) == 0 {
		return "", 0, nil
	}

	maxW := "" // word with max frequency
	maxF := 0  // max frequency of thet word

	for k, v := range m {
		if v > maxF {
			maxW = k
			maxF = v
		}
	}
	return maxW, maxF, nil
}
