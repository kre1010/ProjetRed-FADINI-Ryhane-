package menu

import (
	"fmt"

	"red/src/character"
	"red/src/equipement"
	"red/src/inventory"
	"red/src/sounds"
)

func NouvellePartie() {
	var joueur character.Character

	// FCT DE SAMI
	joueur.MenuInitCharacter()
	// LE PERSO VIENT D'ETRE CREE
	MenuJeu(&joueur)
}

func MenuJeu(player *character.Character) {
	for {
		Nettoyer()

		fmt.Println("\033[35m")
		fmt.Println("========================================")
		fmt.Println("                  KOTW")
		fmt.Println("========================================")
		fmt.Println()
		fmt.Println("1. Jouer")
		fmt.Println("2. Personnage")
		fmt.Println("3. Inventaire")
		fmt.Println("4. Équipement")
		fmt.Println("5. Marchand")
		fmt.Println("6. Armurier")
		fmt.Println("0. Retour")
		fmt.Println()
		fmt.Println("========================================")

		var choix int
		fmt.Print("Choix : ")
		fmt.Scan(&choix)

	Nettoyer()	
		switch choix {
		case 1:
			fmt.Println("Jouer")

		case 2:
			character.ShowCharacterInfo(player)
			fmt.Scanln()
			fmt.Scanln()

		case 3:
			inventory.AccessInventory(player)
			fmt.Scanln()
			fmt.Scanln()
		case 4:
			equipement.AfficherEquipement()
			fmt.Scanln()
			fmt.Scanln()

		case 5:
			inventory.Marchand(player)
			fmt.Scanln()
			fmt.Scanln()

		case 6:
			character.Armurier(player)
			fmt.Scanln()
			fmt.Scanln()

		case 0:
			sounds.StopFlash()
			Menu()
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
