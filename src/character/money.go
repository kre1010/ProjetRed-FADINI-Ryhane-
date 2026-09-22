package character

import (
	"fmt"
)

func gagnerArgent(c *Character, montant int) {
	c.Argent += montant
	fmt.Printf("%s a gagné %d eddie !\n", c.Name, montant)
}

func afficherArgent(c *Character) {
	fmt.Printf("eddie actuel : %d eddie\n", c.Argent)
}

func acheterObjet(c *Character, nomObjet string, prix int) {
	if c.Argent >= prix {
		c.Argent -= prix
		fmt.Printf("Vous avez acheté %s pour %d eddie.\n", nomObjet, prix)
	} else {
		fmt.Printf("Pas assez d'eddie pour acheter %s.\n", nomObjet)
	}
}

func character() {
	character := Character{
		Name:   "character",
		Argent: 100,
	}
	fmt.Println("=== Début du jeu ===")
	afficherArgent(&character)
	fmt.Println("\n=== Combat gagné ===")
	gagnerArgent(&character, 50)
	afficherArgent(&character)
	fmt.Println("\n=== Boutique ===")
	acheterObjet(&character, "canon long", 60)
	afficherArgent(&character)
	acheterObjet(&character, "Medkit", 30)
	afficherArgent(&character)
	acheterObjet(&character, "plaque ceramique", 200)
	afficherArgent(&character)

}
