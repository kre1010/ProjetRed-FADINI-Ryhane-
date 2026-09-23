package character

import (
	"fmt"
)

func ShowCharacterInfo(c *Character) {
	fmt.Println("===========================================")
	fmt.Println("              MON PERSONNAGE")
	fmt.Println("===========================================")
	fmt.Printf("\t nom : %s\n", c.Name)
	fmt.Printf("\t class : %d\n", c.Class)
	fmt.Printf("\t level : %d\n", c.Level)
	fmt.Printf("\t HP : %d/%d\n", c.Hp, c.MaxHp)
	fmt.Printf("\t Eddies : %d\n", c.Argent)
	fmt.Println("===========================================")
	fmt.Println("Appuyer sur entrée pour revenir...")
}


