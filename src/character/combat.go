package character

import "fmt"

func (c *Character, o *Operateur) Fight() {

	for c.Hp > 0 && o.Hp > 0 {
		var choix int
		fmt.Scan(&choix)
		switch choix {
			case 1:
				fmt.Println("attaquer")
		}
	}
}

func combat(c *Character, o *Operateur) {
     damage := c.Attaque
	o.Hp -= damage 
    if o.Hp < 0 {
		fmt.Printf("operateur is dead")
		     return
	
		}else {
	     	damage := o.Attaque
	   		c.Hp -= damage
			if c.Hp < 0 {
			fmt.Printf("you died")
				return
	   }
		      


	


	// 		fmt.Printf("vous attaquer l'operateur %s avec %d de degats\n", o.Nom, c.Attaque)
	// 		fmt.Printf("l'operateur %s vous attaque avec %d de degats\n", o.Nom, o.Attaque)
	// 		return
	// 	} else {
	// 		if c.Hp <= 0 {
	// 			fmt.Println("you died")
	// 		}
	// 		if o.Hp <= 0 {
	// 			fmt.Println("operateur is dead")
	// 			fmt.Println("you win")
	// 		}
	// 	}
	// }