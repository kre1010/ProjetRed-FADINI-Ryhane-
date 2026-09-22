package inventory

import "fmt"

var inventory []string

func AccessInventory() {
	fmt.Println("Inventaire :")

	for i := 0; i < len(inventory); i++ {
		fmt.Println(i+1, "-", inventory[i])
	}
}

func buyItem(item string) {
	if len(inventory) >= 10 {
		fmt.Println("L'inventaire est plein !")
	} else {
		inventory = append(inventory, item)
		fmt.Println(item, "a été ajouté à l'inventaire.")
	}
}



