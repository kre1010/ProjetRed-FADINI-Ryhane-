package jouer

import (
	"fmt"
	"red/src/character"
	"red/src/equipement"
	"red/src/inventory"
	"red/src/sounds"
	"red/src/utils"
	"time"
)

func Intro(c *character.Character) {

	utils.WriteTyper(`Dans un monde où la guerre est devenue un business qui rapporte gros, là où les entreprises privées mènent des guerres sanglantes et sans merci pour des millions d'euros/dollars.

C'est là que le KOTW a été créé, localisé à Kino der Toten dans l'ancienne Allemagne. Plusieurs soldats, des spécialités différentes, des ennemis de plusieurs forces spéciales du monde entier.

Arriverez-vous à trouver votre place, écrire l'histoire en évitant une mort spectaculaire ou renoncer au combat et raconter l'histoire des autres?

`, utils.NormalDelay)

	fmt.Println("Appuyer sur entrée pour continuer...")
	fmt.Scanln()
	fmt.Scanln()

	utils.Nettoyer()

	fmt.Println("========== ENTRAINEMENT ==========")
	fmt.Println()
	fmt.Println("Avant d'entrer dans le KOTW, vous devez prouver votre valeur.")
	fmt.Println("Un agent d'entraînement entre dans l'arène.")
	fmt.Println()

	agent := character.Operateurs["Agent d'entrainement"]

	character.Combat(c, &agent)

	if c.Hp <= 0 {
		fmt.Println()
		fmt.Println("Vous avez perdu l'entraînement.")
		return
	}

	fmt.Println()
	fmt.Println("Entraînement terminé !")
	fmt.Println("Vous êtes maintenant prêt pour le KOTW.")

	time.Sleep(2 * time.Second)

	Preparation(c)

	ennemis := []string{
		"Jaegerkorpset",
		"KSK",
		"Bope",
		"Spetnaz",
		"Neavy Seal",
	}

	for i, nom := range ennemis {

		utils.Nettoyer()

		HistoireQuartier(i)

		if nom == "Bope" {
			EvenementTales(c)

			if c.Hp <= 0 {
				fmt.Println("Votre aventure s'arrête ici.")
				return
			}
		}

		ennemi := character.Operateurs[nom]

		fmt.Println()
		fmt.Println("==============================")
		fmt.Println(c.Name, "VS", ennemi.Nom)
		fmt.Println("==============================")
		fmt.Println()

		character.Combat(c, &ennemi)

		if c.Hp <= 0 {
			fmt.Println()
			fmt.Println("Vous avez perdu contre", ennemi.Nom)
			return
		}

		if ennemi.Hpactuel > 0 {
			fmt.Println("Vous avez quitté le combat.")
			return
		}

		fmt.Println()
		fmt.Println("Vous avez vaincu", ennemi.Nom, "!")

		c.Argent += 75

		fmt.Println()
		fmt.Println("+75 Eddies")
		fmt.Println("Vous possédez maintenant", c.Argent, "Eddies.")

		time.Sleep(2 * time.Second)

		if i < len(ennemis)-1 {

			fmt.Println()
			fmt.Println("Le combat est terminé.")
			fmt.Println("Vous quittez la zone...")

			Preparation(c)
		}
	}
	utils.Nettoyer()

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("         KOTW TERMINÉ")
	fmt.Println("================================")
	sounds.StopMenu()
	fmt.Println()
	fmt.Println("Vous avez vaincu tous les opérateurs !")
	fmt.Println()
	fmt.Println(c.Name, "est le dernier combattant encore debout.")
	fmt.Println()
	fmt.Println("Eddies :", c.Argent)
	fmt.Println("Tales :", c.Inv["Tales"])
}

func Preparation(c *character.Character) {

	for {
		utils.Nettoyer()

		fmt.Println("========== ZONE DE PRÉPARATION ==========")
		fmt.Println()
		fmt.Println("HP :", c.Hp, "/", c.MaxHp)
		fmt.Println("Niveau :", c.Level)
		fmt.Println("Eddies :", c.Argent)
		fmt.Println()
		fmt.Println("1. Aller au marchand")
		fmt.Println("2. Aller à l'armurier")
		fmt.Println("3. Voir mon personnage")
		fmt.Println("4. Voir mon inventaire")
		fmt.Println("5. Voir mon équipement")
		fmt.Println("6. Continuer l'aventure")
		fmt.Println()

		var choix int

		fmt.Print("Choix : ")
		fmt.Scan(&choix)

		switch choix {

		case 1:
			utils.Nettoyer()
			sounds.PlayHein()
			inventory.Marchand(c)

			fmt.Println()
			fmt.Println("Appuyez sur Entrée pour revenir...")
			fmt.Scanln()
			fmt.Scanln()

		case 2:
			utils.Nettoyer()
			character.Armurier(c)

			fmt.Println()
			fmt.Println("Appuyez sur Entrée pour revenir...")
			fmt.Scanln()
			fmt.Scanln()

		case 3:
			utils.Nettoyer()
			character.ShowCharacterInfo(c)

			fmt.Scanln()
			fmt.Scanln()

		case 4:
			utils.Nettoyer()
			inventory.AccessInventory(c)

			fmt.Println()
			fmt.Println("Appuyez sur Entrée pour revenir...")
			fmt.Scanln()
			fmt.Scanln()

		case 5:
			utils.Nettoyer()
			equipement.AfficherEquipement()

			fmt.Scanln()
			fmt.Scanln()

		case 6:
			utils.Nettoyer()
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}



func HistoireQuartier(numero int) {

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("       NOUVEAU QUARTIER")
	fmt.Println("================================")
	fmt.Println()

	switch numero {

	case 0:
		fmt.Println("Vous quittez la zone d'entraînement.")
		fmt.Println("Une lourde porte métallique s'ouvre devant vous.")
		fmt.Println()
		fmt.Println("Vous entrez dans les rues abandonnées autour de Kino der Toten.")
		fmt.Println("Des véhicules militaires brûlés bloquent la route.")
		fmt.Println("Votre premier véritable adversaire vous attend.")

	case 1:
		fmt.Println("Vous quittez le premier quartier.")
		fmt.Println()
		fmt.Println("Vous arrivez dans une ancienne zone industrielle.")
		fmt.Println("Les bâtiments sont abandonnés.")
		fmt.Println("Des impacts de balles couvrent les murs.")
		fmt.Println()
		fmt.Println("Au loin, vous apercevez votre prochain adversaire.")

	case 2:
		fmt.Println("Vous avancez vers un quartier détruit par les combats.")
		fmt.Println()
		fmt.Println("Les rues deviennent de plus en plus étroites.")
		fmt.Println("Une personne vous observe depuis une ruelle.")
		fmt.Println()
		fmt.Println("Vous décidez de vous approcher...")

	case 3:
		fmt.Println("Vous quittez les ruelles et arrivez dans le centre-ville.")
		fmt.Println()
		fmt.Println("Le bruit des combats devient de plus en plus proche.")
		fmt.Println("Il ne reste presque plus personne dans le tournoi.")
		fmt.Println()
		fmt.Println("Un nouvel opérateur bloque votre chemin.")

	case 4:
		fmt.Println("Vous arrivez devant Kino der Toten.")
		fmt.Println()
		fmt.Println("Le quartier est silencieux.")
		fmt.Println("Aucun civil. Aucun soldat.")
		fmt.Println()
		fmt.Println("Une seule personne vous attend.")
		fmt.Println()
		fmt.Println("Le dernier opérateur du KOTW.")
	}

	fmt.Println()
}

func EvenementTales(c *character.Character) {

	fmt.Println("================================")
	fmt.Println("          LA RUELLE")
	fmt.Println("================================")
	fmt.Println()

	fmt.Println("Un homme est assis contre un mur.")
	fmt.Println("Il tient une petite sacoche dans sa main.")
	fmt.Println()
	fmt.Println("Inconnu : Hé, combattant.")
	fmt.Println("Inconnu : J'ai quelque chose qui pourrait t'intéresser.")
	fmt.Println()
	fmt.Println("Vous remarquez plusieurs Tales dans sa sacoche.")
	fmt.Println()

	fmt.Println("Que voulez-vous faire ?")
	fmt.Println()
	fmt.Println("1. Continuer directement vers le prochain combat")
	fmt.Println("2. Parler à l'inconnu")

	var choix int

	fmt.Print("Choix : ")
	fmt.Scan(&choix)

	if choix == 1 {

		fmt.Println()
		fmt.Println("Vous décidez de ne pas perdre de temps.")
		fmt.Println("Vous continuez votre chemin vers le prochain combat.")

		return
	}

	if choix == 2 {

		fmt.Println()
		fmt.Println("Vous vous approchez de l'inconnu.")
		fmt.Println()
		fmt.Println("Inconnu : Ces Tales peuvent être à toi.")
		fmt.Println("Inconnu : Mais seulement si tu réponds correctement à mon énigme.")
		fmt.Println()

		EnigmeTales(c)

		return
	}

	fmt.Println()
	fmt.Println("Vous restez trop longtemps sans répondre.")
	fmt.Println("L'inconnu disparaît dans les ruelles.")
}

func EnigmeTales(c *character.Character) {

	fmt.Println("========== ÉNIGME ==========")
	fmt.Println()
	fmt.Println("Je peux remplir une pièce entière,")
	fmt.Println("mais je ne prends aucune place.")
	fmt.Println()
	fmt.Println("Qui suis-je ?")
	fmt.Println()

	var reponse string

	fmt.Print("Votre réponse : ")
	fmt.Scan(&reponse)

	if reponse == "lumiere" || reponse == "Lumiere" {

		fmt.Println()
		fmt.Println("L'inconnu vous regarde et sourit.")
		fmt.Println()
		fmt.Println("Inconnu : Bonne réponse.")
		fmt.Println()

		c.Inv["Tales"] += 3
		c.Argent += 60

		fmt.Println("Vous recevez 3 Tales et 60 pièces d'or.")
		fmt.Println("Tales :", c.Inv["Tales"])
		fmt.Println("Argent :", c.Argent)

		return
	}

	fmt.Println()
	fmt.Println("Mauvaise réponse.")
	fmt.Println()
	fmt.Println("L'inconnu referme immédiatement sa sacoche.")
	fmt.Println()
	fmt.Println("Inconnu : Dommage.")
	fmt.Println()
	fmt.Println("Il commence à partir avec les Tales.")

	fmt.Println()
	fmt.Println("Que voulez-vous faire ?")
	fmt.Println()
	fmt.Println("1. Le laisser partir")
	fmt.Println("2. Le combattre pour récupérer les Tales")

	var choix int

	fmt.Print("Choix : ")
	fmt.Scan(&choix)

	if choix == 1 {

		fmt.Println()
		fmt.Println("Vous le laissez partir.")
		fmt.Println("Vous reprenez votre chemin.")

		return
	}

	if choix == 2 {

		CombatTales(c)

		return
	}

	fmt.Println("Vous hésitez et l'inconnu disparaît.")
}

func CombatTales(c *character.Character) {

	fmt.Println()
	fmt.Println("Vous sortez votre arme.")
	fmt.Println()
	fmt.Println("Inconnu : Mauvais choix...")
	fmt.Println()

	contrebandier := character.Operateur{
		Nom:      "Contrebandier",
		Hpmax:    60,
		Hpactuel: 60,
		Attaque:  15,
	}

	character.Combat(c, &contrebandier)

	if c.Hp <= 0 {

		fmt.Println()
		fmt.Println("Le contrebandier vous a vaincu.")

		return
	}

	if contrebandier.Hpactuel <= 0 {

		fmt.Println()
		fmt.Println("Vous avez vaincu le contrebandier.")
		fmt.Println()
		fmt.Println("Vous récupérez sa sacoche.")

		c.Inv["Tales"] += 3
		c.Argent += 60

		fmt.Println()
		fmt.Println("Vous recevez 3 Tales et 60 pièces d'or.")
		fmt.Println("Argent :", c.Argent)
		fmt.Println("Tales :", c.Inv["Tales"])
	}
}


