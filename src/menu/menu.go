package menu

import (
	"fmt"
	"os"
	"red/src/sounds"
	"red/src/utils"
)

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

	fmt.Println("\033[1;31m" + `
               /////'
              '  # o
              C   - |
 ___          '  =__'        ___
(\ _ \_       |   |        _/  ')
 \  (__\  ,---- _ |----.  /__)- |
  \__  ( (           /  ) )  __/
    |_X_\/ \.   #  _.|  \/_X_|
      |  \ /(   /    /\ /  |
       \ /  (  ,    /  \ _/
            /______/
           [:::::::]
          /*%*%*%*%*\
          >%*%#%*%*%|
         /%*%*#*%*%*\
        ######^#######
` + "\033[0m")


	sounds.PlayMenu()
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
		sounds.StopMenu()
		utils.Nettoyer()
		fmt.Println("Fermeture de KOTW...")
		os.Exit(0)

	case 1:
		utils.Nettoyer()
		NouvellePartie()

	case 2:
		utils.Nettoyer()
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
		utils.Nettoyer()
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


