package inventory

import (
	"fmt"
	"red/src/character"
)

func AccessInventory(c *character.Character) {
	fmt.Println()
	fmt.Println("==========INVENTAIRE==========")
	fmt.Println()
	fmt.Println("-", c.Inv)
	fmt.Println()
	fmt.Println("Inventaire :", len(c.Inv), "/", c.MaxInv)
	fmt.Println()
	fmt.Println("==============================")
}
