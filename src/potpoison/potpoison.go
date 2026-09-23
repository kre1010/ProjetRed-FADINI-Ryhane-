package potpoison

import (
	"fmt"
	"red/src/character"
	"time"
)

func CocktailMolotov(c *character.Character) {
	for i := 0; i < 3; i++ {
		c.Hp -= 10
		time.Sleep(1 * time.Second)
		fmt.Println("-10 Hp")
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Total : -30 Hp")
}


