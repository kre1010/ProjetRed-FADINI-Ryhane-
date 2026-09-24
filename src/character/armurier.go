package character

import (
	"fmt"
	"red/src/equipement"
)

func Armurier(c *Character) {
	fmt.Println("\033[1;33m=== ARMURIER ===\033[0m")
	fmt.Println("\033[36m1.\033[0m gilet pare-balles militech - 30 eddie")
	fmt.Println("\033[36m2.\033[0m casque pare balles ops-core - 10 eddie")
	fmt.Println("\033[36m3.\033[0m pantalon renforcé crye precision g4 - 10 eddie")
	fmt.Println("\033[36m4.\033[0m gants Oakley SI - 5 eddie")
	fmt.Println("\033[36m5.\033[0m bottes Salomon Forces - 5 eddie")
	fmt.Println("\033[36m6.\033[0m plaque en kevlar - 2 eddie")
	fmt.Println("\033[36m8.\033[0m Augmentation d'inventaire - 30 eddie")
	fmt.Println("\033[36m0.\033[0m Quitter l'armurier")
	fmt.Print("Ton choix : ")

	var choix int
	fmt.Print("Ton choix : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		CreeObjet(c, "Gilet Militech", 5, 3, "Torse")

	case 2:
		CreeObjet(c, "Casque Ops-Core", 3, 2, "Tête")

	case 3:
		CreeObjet(c, "Pantalon Crye Precision G4", 4, 3, "Jambe")

	case 4:
		CreeObjet(c, "Gants Oakley SI", 1, 1, "Main")

	case 5:
		CreeObjet(c, "Bottes Salomon Forces", 2, 2, "Pieds")

	case 0:
		fmt.Println("Tu quittes l'armurier.")

	default:
		fmt.Println("Choix invalide.")
	}
}

func CreeObjet(c *Character, nomObjet string, kevlar int, ceramic int, emplacement string) {
	if c.Inv["Kevlar"] < kevlar {
		fmt.Println("Pas assez de Kevlar.")
		return
	}

	if c.Inv["Céramic"] < ceramic {
		fmt.Println("Pas assez de Céramic.")
		return
	}

	c.Inv["Kevlar"] -= kevlar
	c.Inv["Céramic"] -= ceramic

	switch emplacement {
	case "Tête":
		equipement.Equip.Tête[nomObjet]++

	case "Torse":
		equipement.Equip.Torse[nomObjet]++

	case "Jambe":
		equipement.Equip.Jambe[nomObjet]++

	case "Main":
		equipement.Equip.Mains[nomObjet]++

	case "Pieds":
		equipement.Equip.Pieds[nomObjet]++
	}

	fmt.Println("Vous avez fabriqué", nomObjet)
}
