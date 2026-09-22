package inventory

import (
	"fmt"
	"red/src/character"
)

func Marchand(c *character.Character) {
	fmt.Println("0. Quitter")
	fmt.Println("1. Gilet par-balle")
	fmt.Println("2. Canon Long")
	fmt.Println("3. Bandage")
	fmt.Println("4. Heavy Bullets")
	fmt.Println("5. Medium Bullets")
	fmt.Println("6. Shotgun shells")
	fmt.Println("7. Potion de soin")


	var choix int
	fmt.Print("\nChoix : ")
	fmt.Scan(&choix)

	switch choix {
	case 0:
		fmt.Println("for")
	case 1:
		fmt.Println("Vous avez acheté x1 Gilet par-balle !")
		c.Inv = map[string]int{"Gilet par-balle" : 1} 
	case 2:
		fmt.Println("Vous avez actheter x1 Canon Long !")
		c.Inv = map[string]int{"Canon Long" : 1}
	case 3:
		fmt.Println("Vous avez acheter x5 Bandage !")
		c.Inv = map[string]int{"Bandage" : 5}
	case 4:
		fmt.Println("Vous avez acheter x3 Heavy Bullets !")
		c.Inv = map[string]int{"Heavy Bullets" : 3}
	case 5:
		fmt.Println("Vous avez acheté x15 Medium bullets !")
		c.Inv = map[string]int{"Medium Bullets" : 15}
	case 6:
		fmt.Println("Vous avez actheter x4 Shotgun shells !")
		c.Inv = map[string]int{"Shotgun shells" : 1}
	case 7:
		fmt.Println("Vous avez réçu une Potion de soin gratuite !")
		c.Inv = map[string]int{"Potion de soin" : 1}
	default:
		fmt.Println("Choix invalide.")
	}
}
