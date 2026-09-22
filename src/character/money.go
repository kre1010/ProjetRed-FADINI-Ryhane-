package character

import (
	"fmt"
)

func gagnerArgent(j *Character, montant int) {
	j.Argent += montant
	fmt.Printf("%s a gagné %d eddie !\n", j.Name, montant)
}

func afficherArgent(j *Character) {
	fmt.Printf("eddie actuel : %d eddie\n", j.Argent)
}

func acheterObjet(j *Character, nomObjet string, prix int) {
	if j.Argent >= prix {
		j.Argent -= prix
		fmt.Printf("Vous avez acheté %s pour %d eddie.\n", nomObjet, prix)
	} else {
		fmt.Printf("Pas assez d'eddie pour acheter %s.\n", nomObjet)
	}
}

func main() {
	character := Character{
		Name:   "character",
		Argent: 100,
	}
	fmt.Println("=== Début du jeu ===")
	afficherArgent(&character)
	fmt.Println("\n=== Combat gagné ===")
	gagnerArgent(&character, 100)
	afficherArgent(&character)
	fmt.Println("\n=== Boutique ===")
	acheterObjet(&character, "canon long", 60)
	afficherArgent(&character)
	acheterObjet(&character, "Medkit", 30)
	afficherArgent(&character)
	acheterObjet(&character, "plaque ceramique", 200)
	afficherArgent(&character)

}
