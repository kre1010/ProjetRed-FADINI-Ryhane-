package character

import (
	"fmt"
)

func (c *Character) ShowCharacterInfo() {
	fmt.Println("Voici les informations de votre personnage:")
	fmt.Printf("\t nom : %s\n", c.Name)
	fmt.Printf("\t class : %d\n", c.Class)
	fmt.Printf("\t level : %s\n", c.Level)
	fmt.Printf("\t HP : %d/%d\n", c.Hp,c.maxHp)
}
