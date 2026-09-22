package inventory

import "fmt"

func AccessInventory(inventory []string) {
	fmt.Println("Inventaire :")

	for _, item := range inventory {
		fmt.Println(item)
	}
}