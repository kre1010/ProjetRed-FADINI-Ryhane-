package inventory

import (
	"fmt"

	"red/src/character"
)

func Marchand(c *character.Character) {

	AfficherMarchand(c)

	for {
		var choix int

		fmt.Print("Choix : ")
		fmt.Scan(&choix)

		switch choix {

		case 0:
			fmt.Println("Vous quittez le marchand.")
			return

		case 1:
			if c.AcheterObjet("Kevlar", 8) {
				c.Inv["Kevlar"]++
			}

		case 2:
			if c.AcheterObjet("Canon Long", 20) {
				c.Inv["Canon Long"]++
			}

		case 3:
			if c.AcheterObjet("Bandage", 10) {
				c.Inv["Bandage"]++
			}

		case 4:
			if c.AcheterObjet("Heavy Bullets", 15) {
				c.Inv["Heavy Bullets"]++
			}

		case 5:
			if c.AcheterObjet("Medium Bullets", 15) {
				c.Inv["Medium Bullets"]++
			}

		case 6:
			if c.AcheterObjet("Shotgun Shells", 15) {
				c.Inv["Shotgun Shells"]++
			}

		case 7:
			if c.Inv["Potion de Soin"] > 0 {
				fmt.Println("Vous avez déjà récupéré la Potion de Soin gratuite.")
			} else {
				c.Inv["Potion de Soin"]++
				fmt.Println("Vous recevez une Potion de Soin gratuitement !")
			}
		case 8:
			if c.AcheterObjet("Céramic", 10) {
				c.Inv["Céramic"]++
			}

		case 9:
			if c.AcheterObjet("Grenade", 25) {
				c.Inv["Grenade"]++
			}

		case 10:
			if c.AcheterObjet("MedKit", 15) {
				c.Inv["MedKit"]++
			}

		case 11:
			if c.AcheterObjet("Cocktail Molotov", 20) {
				c.Inv["Cocktail Molotov"]++
			}
		
		case 12:
			if c.AcheterObjet("Augmentation d'Inventaire", 50) {
				c.MaxInv += 10
				fmt.Println("Votre inventaire a été augmenté de 10 places !")
			}

		default:
			fmt.Println("Choix invalide.")
		}
		fmt.Printf("Eddies restants : %d\n\n", c.Argent)
	}
}

func AfficherMarchand(c *character.Character) {
	fmt.Println("========== MARCHAND ==========")
	fmt.Println("------------------------------")
	fmt.Println("0. Quitter")
	fmt.Println("1. Kevlar : 8 Eddies")
	fmt.Println("2. Canon Long : 20 Eddies")
	fmt.Println("3. Bandage : 10 Eddies")
	fmt.Println("4. Heavy Bullets : 15 Eddies")
	fmt.Println("5. Medium Bullets : 15 Eddies")
	fmt.Println("6. Shotgun Shells : 15 Eddies")
	fmt.Println("7. Potion de soin : GRATUIT !")
	fmt.Println("8. Céramic : 10 Eddies")
	fmt.Println("9. Grenade : 25 Eddies")
	fmt.Println("10. MedKit : 15 Eddies")
	fmt.Println("11. Cocktail Molotov : 20 Eddies")
	fmt.Println("12. Augmentation d'Inventaire : 50 Eddies")
	fmt.Println("==============================")
}




