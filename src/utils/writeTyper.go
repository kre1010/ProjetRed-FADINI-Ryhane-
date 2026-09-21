package utils

import (
	"fmt"
	"time"
)

const (
	NormalDelay = 30
)

func WriteTyper(str string, delay int) {
	for _, char := range str {
		fmt.Print(char)
		time.Sleep(time.Duration(delay) * time.Microsecond)
	}
}
