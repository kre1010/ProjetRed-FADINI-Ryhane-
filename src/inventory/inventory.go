package inventory

import "fmt"

var inventory map[string]int

func AccessInventory() {
	fmt.Println("Inventaire :")
	fmt.Println("-", inventory)
}
