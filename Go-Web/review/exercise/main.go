package main

import (
	"exercise/exercise"
	"fmt"
)

func main() {
	item1 := exercise.Item{Name: "Sword", Type: "Weapon"}
	item2 := exercise.Item{Name: "Health Potion", Type: "Consumable"}
	item3 := exercise.Item{Name: "Shield", Type: "Armor"}

	player1 := exercise.Player{Name: "Hero"}
	player1.PickItem(item1)
	player1.PickItem(item2)
	fmt.Println(player1.UseItem("Sword"))

	player2 := exercise.Player{Name: "Warrior"}
	player2.PickItem(item3)
	fmt.Println(player2.UseItem("Health Potion"))

}
