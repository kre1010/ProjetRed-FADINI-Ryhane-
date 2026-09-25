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
			if c.AcheterObjet("Heavy Bullets", 4) {
				c.Inv["Heavy Bullets"]++
			}

		case 4:
			if c.AcheterObjet("Medium Bullets", 3) {
				c.Inv["Medium Bullets"]++
			}

		case 5:
			if c.AcheterObjet("Shotgun Shells", 2) {
				c.Inv["Shotgun Shells"]++
			}

		case 6:
			if c.Inv["Bandage"] > 0 {
				fmt.Println("Vous avez déjà récupéré la Potion de Soin gratuite.")
			} else {
				c.Inv["Bandage"]++
				fmt.Println("Vous recevez une Potion de Soin gratuitement !")
			}
		case 7:
			if c.AcheterObjet("Céramic", 10) {
				c.Inv["Céramic"]++
			}

		case 8:
			if c.AcheterObjet("Grenade", 30) {
				c.Inv["Grenade"]++
			}

		case 9:
			if c.AcheterObjet("MedKit", 25) {
				c.Inv["MedKit"]++
			}

		case 10:
			if c.AcheterObjet("Cocktail Molotov", 20) {
				c.Inv["Cocktail Molotov"]++
			}
		
		case 11:
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
			marchand := `
         ______________________________________
        /                                      \
       |       top 10 marchand all time         |
        \______________________________________/
              ||                        ||
    .---------||------------------------||---------.
   /  __________________________________________  \
  |  /                                          \  |
  | |                 achète                     | |
  |  \__________________________________________/  |
   \______________________________________________/
         ||                            ||
         ||         .--------.         ||
         ||        /   ____   \        ||
         ||       |   / o  o \ |       ||
         ||       |  |   --   ||       ||
         ||       |   \______/ |       ||
         ||      _/\____|__|__/\_      ||
         ||     / /|           | \     ||
         ||    / / |___________|  \    ||
         ||   L_L  |   [   ]   |   L_L ||
  =======||========|___[___]___|=======||=======
  |           [      ]   [    ]   [   ]           |
  |   ( ) )     (iii)    /===    ( * )    (o)     |
  |_______________________________________________|
       |__|                                 |__|
      ( - )                                ( - )
`
	fmt.Print(marchand)
	fmt.Println("========== MARCHAND ==========")
	fmt.Println("------------------------------")
	fmt.Println("0. Quitter")
	fmt.Println("1. Kevlar : 8 Eddies")
	fmt.Println("2. Canon Long : 20 Eddies")
	fmt.Println("3. Heavy Bullets : 4 Eddies")
	fmt.Println("4. Medium Bullets : 3 Eddies")
	fmt.Println("5. Shotgun Shells : 2 Eddies")
	fmt.Println("6. Potion de soin : GRATUIT !")
	fmt.Println("7. Céramic : 10 Eddies")
	fmt.Println("8. Grenade : 30 Eddies")
	fmt.Println("9. MedKit : 25 Eddies")
	fmt.Println("10. Cocktail Molotov : 20 Eddies")
	fmt.Println("11. Augmentation d'Inventaire : 50 Eddies")
	fmt.Println("==============================")
}





