package menu

import (
	"fmt"
	"os"
	"os/exec"
)

func Nettoyer() {
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func Menu() {
	fmt.Println(`
		██╗  ██╗   ██████╗  ████████╗ ██╗    ██╗
		██║ ██╔╝  ██╔═══██╗ ╚══██╔══╝ ██║    ██║
		█████╔╝   ██║   ██║    ██║    ██║ █╗ ██║
		██╔═██╗   ██║   ██║    ██║    ██║███╗██║
		██║  ██╗  ╚██████╔╝    ██║    ╚███╔███╔╝
		╚═╝  ╚═╝   ╚═════╝     ╚═╝     ╚══╝╚══╝

	                    	KOTW
`)

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
		os.Exit(0)

	case 1:
		Nettoyer()
		NouvellePartie()

	case 2:
		Nettoyer()
		fmt.Println("============== OPTIONS ==============")
		fmt.Println()
		fmt.Println("1. Son")
		fmt.Println("2. Graphismes")
		fmt.Println("3. Retour")
		fmt.Println()
		fmt.Println("Appuie sur Entrée pour revenir...")
		fmt.Scanln()
		fmt.Scanln()

	case 3:
		Nettoyer()
		fmt.Println("=============== CRÉDIT ===============")
		fmt.Println()
		fmt.Println("KOTW")
		fmt.Println("Créé en Go")
		fmt.Println()
		fmt.Println("Appuie sur Entrée pour revenir...")
		fmt.Scanln()
		fmt.Scanln()

	default:
		fmt.Println("Choix invalide.")
	}
}


