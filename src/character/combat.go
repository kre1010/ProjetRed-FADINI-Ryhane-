package character

func (c *Character, o *Operateur) Fight() {
      
	
	
	func (c *character) IsAlive() bool {
		    return c.Hp > 0 




	if c.Hp > 0 && o.Hp > 0 {



   fmt.Printf("vous attaquer l'operateur %s avec %d de degats\n", o.Nom, c.Attaque) 
   fmt.Printf("l'operateur %s vous attaque avec %d de degats\n", o.Nom, o.Attaque)
         return
    } else{
		if c.Hp <= 0 {
			fmt.Println("you died")
		}
		if o.Hp <= 0 {
			fmt.Println("operateur is dead")
			fmt.Println("you win")
		}
	}
}


}
