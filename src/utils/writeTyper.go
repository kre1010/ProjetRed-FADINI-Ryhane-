package utils

import (
	"fmt"
	"time"
)

const (
	NormalDelay = 30
)

func WriteTyper(str string, delay int) {
	fmt.Println("\033[37m")
	for _, char := range str {
		fmt.Print(string(char))
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}
