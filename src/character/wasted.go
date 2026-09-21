package character

import "fmt"

<<<<<<< HEAD
type character struct {
	Name   string
	Health int
	Alive  bool
}

func player() {
	player := character{}
	damage(&player, 30)
	damage(&player, 50)
	damage(&player, 25)
}

func damage(p *character, amount int) {
	if !p.Alive {
		return
	}
	p.Health -= amount
	if p.Health <= 0 {
		p.Health = 0
		p.Alive = false
		fmt.Println(p.Name, "You Died")
	} else {
		fmt.Println(p.Name, "a", p.Health, "HP")
	}
}

=======
func Damage(c *Character, amount int) {

	if c.Hp <= 0 {
		fmt.Println(c.Name, "is already dead.")
		return
	}
	c.Hp -= amount
	if c.Hp <= 0 {
		c.Hp = 0

		fmt.Println(c.Name, "you died !")
	} else {
		fmt.Printf("%s has %d HP remaining.\n", c.Name, c.Hp)
	}
}
>>>>>>> c5c2a9623efdfeb4b1b13f4df7fba70ce85f5624
