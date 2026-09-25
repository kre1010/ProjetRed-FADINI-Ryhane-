package utils

import (
	"os/exec"
	"os"
)

func Nettoyer() {

    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}
