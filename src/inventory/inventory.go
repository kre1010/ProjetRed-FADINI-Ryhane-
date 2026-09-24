package inventory

import (
	"fmt"
	"red/src/character"
)

func AccessInventory(c *character.Character) {
	fmt.Println("Inventaire :")
	fmt.Println("-", c.Inv)
	fmt.Println()
	fmt.Println("Inventaire :", len(c.Inv), "/", c.MaxInv)
}
