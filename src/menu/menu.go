package menu

import (
	"fmt"
	"os"
	"os/exec"
	"red/src/sounds"
)

func Nettoyer() {
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func Menu() {	
	fmt.Println("\033[31m")
	fmt.Println(`
		██╗  ██╗   ██████╗  ████████╗ ██╗    ██╗
		██║ ██╔╝  ██╔═══██╗ ╚══██╔══╝ ██║    ██║
		█████╔╝   ██║   ██║    ██║    ██║ █╗ ██║
		██╔═██╗   ██║   ██║    ██║    ██║███╗██║
		██║  ██╗  ╚██████╔╝    ██║    ╚███╔███╔╝
		╚═╝  ╚═╝   ╚═════╝     ╚═╝     ╚══╝╚══╝

	                    	KOTW
`)

	fmt.Println("\033[33m")
	fmt.Println("0. Quitter")
	fmt.Println("1. Nouvelle Partie")
	fmt.Println("2. Options")
	fmt.Println("3. Crédit")

	var choix int
	fmt.Print("\nChoix : ")
	fmt.Scan(&choix)

	switch choix {
	case 0:
		Nettoyer()
		fmt.Println("Fermeture de KOTW...")
		sounds.StopFlash()
		os.Exit(0)

	case 1:
		Nettoyer()
		NouvellePartie()

	case 2:
		Nettoyer()
		fmt.Println("\033[32m")
		fmt.Println("============== OPTIONS ==============")
		fmt.Println()
		fmt.Println("1. Son")
		fmt.Println("2. Graphismes")
		fmt.Println("3. Retour")
		fmt.Println()
		fmt.Println("Appuie sur Entrée pour revenir...")
		fmt.Scanln()
		fmt.Scanln()

		Menu()
		

	case 3:
		Nettoyer()
		fmt.Println("\033[32m")
		fmt.Println("=============== CRÉDIT ===============")
		fmt.Println()
		fmt.Println("KOTW")
		fmt.Println("Créé en Go par Sami ABIROU, Ryhane FADINI et Killian MORETTE")
		fmt.Println("Avec l'aide de CYRIL, LILIAN et VITO")
		fmt.Println()
		fmt.Println("======================================")
		fmt.Println()
		fmt.Println("Appuie sur Entrée pour revenir au menu...")
		fmt.Scanln()
		fmt.Scanln()

		Menu()

	default:
		fmt.Println("Choix invalide.")
	}
}


