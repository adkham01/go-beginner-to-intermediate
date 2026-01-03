package main

import (
	"exercise1/exercise1"
	"fmt"
)

func main() {
	item1 := exercise1.Item{Name: "Sword", Type: "Weapon"}
	item2 := exercise1.Item{Name: "Health Potion", Type: "Consumable"}
	item3 := exercise1.Item{Name: "Shield", Type: "Armor"}

	player1 := exercise1.Player{Name: "Hero"}
	player1.PickItem(item1)
	player1.PickItem(item2)
	fmt.Println(player1.UseItem("Sword"))

	player2 := exercise1.Player{Name: "Warrior"}
	player2.PickItem(item3)
	fmt.Println(player2.UseItem("Health Potion"))

}
