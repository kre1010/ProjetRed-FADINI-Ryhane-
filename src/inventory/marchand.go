package inventory

import (
	"fmt"
	"red/src/character"
	// "red/src/menu"
)

func Marchand(c *character.Character) {
	fmt.Println("0. Quitter")
	fmt.Println("1. Kevlar : 8 Eddies")
	fmt.Println("2. Canon Long : 20 Eddies")
	fmt.Println("3. Bandage : 10 Eddies")
	fmt.Println("4. Heavy Bullets : 15 Eddies")
	fmt.Println("5. Medium Bullets : 15 Eddies")
	fmt.Println("6. Shotgun shells : 15 Eddies")
	fmt.Println("7. Potion de soin : GRATUIT !")
	fmt.Println("8. Sac : 5 Eddies")
	fmt.Println("9. Céramic : 10 Eddies")

	var choix int
	fmt.Print("\nChoix : ")
	fmt.Scan(&choix)

	switch choix {
	case 0:
		fmt.Println("alo")
	case 1:
		c.AcheterObjet("Kevlar", 8)
	case 2:
		c.AcheterObjet("Canon Long", 20)
	case 3:
		c.AcheterObjet("Bandage", 10)
	case 4:
		c.AcheterObjet("Heavy Bullets", 15)
	case 5:
		c.AcheterObjet("Medium Bullets", 15)
	case 6:
		c.AcheterObjet("Shotgun Shells", 15)
	case 7:
		c.AcheterObjet("Potion de Soin", 0)
	case 8:
		c.AcheterObjet("Sac", 5)
	case 9:
		c.AcheterObjet("Céramic", 10)
	default:
		fmt.Println("Choix invalide.")
	}
}
