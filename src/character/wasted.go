package character

import "fmt"

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
