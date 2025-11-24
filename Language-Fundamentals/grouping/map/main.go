package main

import "fmt"

func main() {
	m := map[string]int{
		"Todd":  42,
		"Henry": 16,
		"Paget": 14,
	}

	// --- or ---
	// m := make(map[string] int)
	// m["Todd"] = 42
	// m["Henry"] = 162

	fmt.Println("Henry's age is", m["Henry"])

	for k := range m {
		fmt.Printf("Just keys: %s\n", k)
	}
	fmt.Println(m)
	fmt.Printf("%#v\n", m)

	m["George"] = 78

	for k, v := range m {
		fmt.Println(k, v)
	}
	for _, v := range m {
		fmt.Println(v)
	}

	delete(m, "Paget")
	fmt.Println(m)

	v, ok := m["Paget"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key does not exist")
	}

	if v, ok := m["Lukas"]; ok {
		fmt.Println("The value prins", v)
	} else {
		fmt.Println("Key does not exist")
	}

}
