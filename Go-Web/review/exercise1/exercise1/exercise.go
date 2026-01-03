package exercise1

import "fmt"

type Player struct {
	Name      string
	Inventory []Item
}

type Item struct {
	Name string
	Type string
}

func (p *Player) PickItem(item Item) {
	p.Inventory = append(p.Inventory, Item{Name: item.Name, Type: item.Type})
	fmt.Println(p.Name + " picked up a " + item.Name + " (" + item.Type + ")")
}

func (p *Player) DropItem(item string) {
	for i, it := range p.Inventory {
		if it.Name == item {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Println(p.Name + " dropped the " + item)
			return
		}
	}
}

func (p *Player) UseItem(item string) string {
	for _, it := range p.Inventory {
		if it.Name == item {
			fmt.Println(p.Name + " is using the " + it.Name)
			return p.Name + " uses the " + it.Name + " (" + it.Type + ")"
		}
	}
	return p.Name + " does not have the " + item
}
