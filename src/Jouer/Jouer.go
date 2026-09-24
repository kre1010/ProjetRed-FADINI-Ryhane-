package jouer

import (
	"fmt"
	"red/src/utils"
)

func Intro() {
	utils.WriteTyper(`Dans un monde où la guerre est devenue un business qui rapporte gros, là où les entreprises privées mènent des guerres sanglantes et sans merci pour des millions d'euros/dollars.

C'est là que le KOTW a été créé, localisé à Kino der Toten dans l'ancienne Allemagne. Plusieurs soldats, des spécialités différentes, des ennemis de plusieurs forces spéciales du monde entier.

Arriverez-vous à trouver votre place, écrire l'histoire en évitant une mort spectaculaire ou

renoncer au combat et raconter l'histoire des autres?

`, utils.NormalDelay)

	fmt.Println("Appuyer sur entrée pour continuer...")
	fmt.Scanln()
	fmt.Scanln()
	utils.Nettoyer()
	Training()

}

func Training(){
//	initOperateur(operateur["Agent d'entrainement"])
	fmt.Println("Bonjour !")

}


