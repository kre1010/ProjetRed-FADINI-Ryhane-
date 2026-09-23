package menu

import (
	"fmt"

	"red/src/character"
	"red/src/equipement"
	"red/src/inventory"
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

		var choix int
		fmt.Print("Choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			fmt.Println("Jouer")

		case 2:
			Nettoyer()
			character.ShowCharacterInfo(player)
			fmt.Scanln()
			fmt.Scanln()

		case 3:
			Nettoyer()
			inventory.AccessInventory(player)
			fmt.Scanln()
			fmt.Scanln()
		case 4:
			Nettoyer()
			equipement.AfficherEquipement()
			fmt.Scanln()
			fmt.Scanln()

		case 5:
			Nettoyer()
			inventory.AfficherMarchand(player)
			inventory.Marchand(player)
			fmt.Scanln()
			fmt.Scanln()

		case 6:
			Nettoyer()
			character.Armurier(player)
			fmt.Scanln()
			fmt.Scanln()

		case 0:
			Nettoyer()
			Menu()
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
