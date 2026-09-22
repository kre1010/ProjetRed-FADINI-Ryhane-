package inventory

import (
	"fmt"
	"red/src/character"
)

func AccessInventory(c *character.Character) {
	fmt.Println("Inventaire :")
	fmt.Println("-", c.Inv)
}

