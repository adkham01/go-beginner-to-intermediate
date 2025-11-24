package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []User
type ByName []User

func (ba ByAge) Len() int {
	return len(ba)
}

func (ba ByAge) Swap(i, j int) {
	ba[i], ba[j] = ba[j], ba[i]
}

func (ba ByAge) Less(i, j int) bool {
	return ba[i].Age < ba[j].Age
}

func (bn ByName) Len() int {
	return len(bn)
}

func (bn ByName) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}

func (bn ByName) Less(i, j int) bool {
	return bn[i].First < bn[j].First
}

func main() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
		},
	}

	u2 := User{
		First: "Jane",
		Last:  "Tomas",
		Age:   35,
		Sayings: []string{
			"Drop by drop bucket gets full",
			"Persistently patiently you are bounce to succeed",
		},
	}

	u3 := User{
		First: "Oliver",
		Last:  "Girud",
		Age:   35,
		Sayings: []string{
			"Success cames by sacrifise",
			"Start from NOW!",
		},
	}

	users := []User{u1, u2, u3}

	fmt.Println("=== ORIGINAL USERS ===")
	fmt.Println(users)

	fmt.Println("\n=== JSON OUTPUT ===")
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong", err)
	}

	fmt.Println("\n=== USERS BEFORE SORTING ===")
	printUsers(users)

	fmt.Println("\n=== USERS AFTER SORTING BY AGE ===")
	byAge := ByAge(users)
	sort.Sort(byAge)
	printUsers(users)

	fmt.Println("\n=== USERS AFTER SORTING BY NAME ===")
	byName := ByName(users)
	sort.Sort(byName)
	printUsers(users)
}

func printUsers(users []User) {
	for i, user := range users {
		fmt.Printf("%d. %s %s (Age: %d)\n", i+1, user.First, user.Last, user.Age)
		fmt.Println("   Sayings:")
		for j, saying := range user.Sayings {
			fmt.Printf("     %d. \"%s\"\n", j+1, saying)
		}
		fmt.Println()
	}
}
