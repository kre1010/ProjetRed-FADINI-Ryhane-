package sounds

import (
	"os/exec"
)

var soundProcess *exec.Cmd

func PlayFlash() {
	soundProcess = exec.Command(
		"powershell",
		"-NoProfile",
		"-Command",
		`$player = New-Object System.Media.SoundPlayer ".\IMG_0075.wav"; $player.PlaySync()`,
	)

	soundProcess.Start()
}

func StopFlash() {
	if soundProcess != nil && soundProcess.Process != nil {
		soundProcess.Process.Kill()
		soundProcess = nil
	}
}
