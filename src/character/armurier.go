package character

import (
	"fmt"
	"red/src/equipement"
)

func Armurier(c *Character) {
	fmt.Println("========== ARMURIER ==========")
	fmt.Println("1. gilet pare-balles militech - 5 Kevlar, 3 céramic")
	fmt.Println("2. casque pare balles ops-core - 3 Kevlar, 2 céramic")
	fmt.Println("3. pantalon renforcé crye precision g4 - 4 Kevlar, 3 céramic ")
	fmt.Println("4. gants Oakley SI - 1 Kevlar, 1 céramic")
	fmt.Println("5. 0m bottes Salomon Forces - 2 Kevlar, 2 céramic")
	fmt.Println("6. Augmentation d'inventaire - 50 eddie")
	fmt.Println("0. Quitter l'armurier")

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
	
	case 6:
		c.AcheterObjet("Augmentation d'inventaire", 50)

	case 0:
		fmt.Println("Tu quittes l'armurerie.")

	default:
		fmt.Println("Choix invalide.")
	}
}


func CreeObjet(c *Character, nomObjet string, kevlar int, ceramic int, emplacement string) {
	if equipement.Equip.Tête[nomObjet] == 1 {
		fmt.Println("Vous possédez déjà l'équipement")
		return
	}

	if equipement.Equip.Torse[nomObjet] == 1 {
		fmt.Println("Vous possédez déjà l'équipement")
		return
	}

	if equipement.Equip.Jambe[nomObjet] == 1 {
		fmt.Println("Vous possédez déjà l'équipement")
		return
	}

	if equipement.Equip.Mains[nomObjet] == 1 {
		fmt.Println("Vous possédez déjà l'équipement")
		return
	}

	if equipement.Equip.Pieds[nomObjet] == 1 {
		fmt.Println("Vous possédez déjà l'équipement")
		return
	}

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
