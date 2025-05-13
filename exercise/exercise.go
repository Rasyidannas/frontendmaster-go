package exercise

import (
	"fmt"
	// "reflect"
)

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name string
	Inventory []Item
}

func(p *Player) PickUpItem(item Item){
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s as picked up %s!\n", p.Name, item.Name)
}

func(p *Player) DropItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...) // Remove item
			fmt.Printf("%s dropped %s,\n", p.Name, itemName)
			return
		}
	}

	fmt.Printf("%s does not have %s in inventory,\n", p.Name, itemName)
}

func(p *Player) UseItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			if item.Type == "potion" {
				fmt.Printf("%s used %s and feels rejuvenated!\n", p.Name, itemName)
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1]...) // remove the potion
			} else {
				fmt.Printf("%s used %s,\n", p.Name, itemName)
			}
			return
		}
	}
}

// func main() {
// 	fmt.Println("exercise executed")
//
// 	player := Player{Name: "Ciko", Inventory: []string{"Blade", "Gun"}}
// 	fmt.Println(player)
// 	fmt.Println(reflect.TypeOf(player.Inventory))
//
// 	newItem := []string{"medkit", "ammo"}
// 	player.PickUpItem(newItem)
// 	fmt.Println(player)
//
// 	player.UseItem("Gun")
//
// 	player.DropItem()
// 	fmt.Println(player)
// }
