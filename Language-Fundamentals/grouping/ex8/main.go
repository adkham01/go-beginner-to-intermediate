package main

import "fmt"

func main() {
	m := make(map[string]int)

	xs := []string{
		"Across", "the", "valley,", "the", "lights", "of", "the", "city", "glittered", "like", "a", "scattered", "handful",
		"of", "diamonds.", "From", "this", "distance,", "the", "chaos", "and", "noise", "resolved", "into", "a", "serene,",
		"silent", "tapestry", "of", "illumination.", "Each", "tiny", "point", "of", "light", "represented", "a", "life,",
		"a", "story,", "a", "collection", "of", "hopes", "and", "dreams.", "He", "often", "wondered", "how", "many", "of",
		"those", "stories", "were", "happy,", "how", "many", "were", "tinged", "with", "the", "same", "melancholy", "that",
		"settled", "in", "his", "own", "chest.", "The", "cool", "night", "air", "carried", "the", "scent", "of", "pine", "and",
		"damp", "earth,", "a", "familiar", "comfort", "that", "did", "little", "to", "ease", "the", "restlessness", "in",
		"his", "soul.", "Tomorrow", "would", "bring", "new", "challenges,", "new", "decisions,", "but", "for", "now",
		"there", "was", "only", "this", "quiet", "moment", "of", "observation.",
		"The", "quantum", "resonance", "fluctuated", "beyond", "predicted", "parameters,", "causing", "the",
		"stabilizers", "to", "engage", "with", "a", "series", "of", "audible", "clicks.", "Dr.", "Aris", "noted", "the",
		"anomaly", "in", "her", "log,", "her", "stylus", "moving", "swiftly", "across", "the", "digital", "pad.", "The",
		"phenomenon", "was", "both", "beautiful", "and", "terrifying,", "a", "shimmering", "cascade", "of", "energy",
		"that", "defied", "conventional", "physics.", "If", "her", "calculations", "were", "correct,", "they",
		"were", "on", "the", "verge", "of", "a", "breakthrough", "that", "would", "redefine", "human", "understanding", "of",
		"dimensional", "theory.", "The", "potential", "applications", "were", "staggering,", "but", "so",
		"were", "the", "risks.", "One", "misstep", "could", "unravel", "the", "very", "fabric", "of", "their", "reality,",
		"a", "responsibility", "that", "weighed", "heavily", "on", "her", "shoulders", "with", "every", "passing",
		"second.",
	}

	for _, v := range xs {
		m[v]++
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
