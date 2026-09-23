package character

import (
	"fmt"
)

func (c *Character) ShowCharacterInfo() {
	fmt.Println("Voici les informations de votre personnage:")
	fmt.Printf("\t nom : %s\n", c.Name)
	fmt.Printf("\t class : %s\n", c.Class)
	fmt.Printf("\t level : %d\n", c.Level)
	fmt.Printf("\t HP : %d/%d\n", c.Hp, c.MaxHp)
	for _, s := range c.skill {
		fmt.Println("Skill", s)
	}
}
