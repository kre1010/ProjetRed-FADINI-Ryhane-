package menu

import (
	"fmt"

	"red/src/character"
	"red/src/inventory"
)

func NouvellePartie() {
    var player character.Character

    // FCT DE SAMI
    player.MenuInitCharacter()
	// LE PERSO VIENT D'ETRE CREE
    MenuJeu(&player)
}

func MenuJeu(player *character.Character) {
    for {
        nettoyer()

        fmt.Println("========================================")
        fmt.Println("                 KOTW")
        fmt.Println("========================================")
        fmt.Println()
        fmt.Println("1. Jouer")
        fmt.Println("2. Personnage")
        fmt.Println("3. Inventaire")
        fmt.Println("4. Équipement")
        fmt.Println("0. Retour")
        fmt.Println()

        var choix int
        fmt.Print("Choix : ")
        fmt.Scan(&choix)

        switch choix {
        case 1:
            // PLUS TARD
            fmt.Println("Jouer")

        case 2:
            // PLUS TARD
            fmt.Println("Personnage")

        case 3:

            inventory.AccessInventory()

        case 4:
            // PLUS TARD
            fmt.Println("Équipement")

        case 0:
            Menu()

        default:
            fmt.Println("Choix invalide.")
        }

        fmt.Println()
        fmt.Println("Appuie sur Entrée pour continuer...")
        fmt.Scanln()
        fmt.Scanln()
    }
}
