package sounds

import (
	"os/exec"
)

var soundProcess *exec.Cmd

func PlayMenu() {
	soundProcess = exec.Command(
		"powershell",
		"-NoProfile",
		"-Command",
		`$player = New-Object System.Media.SoundPlayer ".\IMG_0077.wav"; $player.PlaySync()`,
	)

	soundProcess.Start()
}

func PlayCombat() {
	soundProcess = exec.Command(
		"powershell",
		"-NoProfile",
		"-Command",
		`$player = New-Object System.Media.SoundPlayer ".\IMG_0078.wav"; $player.PlaySync()`,
	)

	soundProcess.Start()
}

func StopMenu() {
	if soundProcess != nil && soundProcess.Process != nil {
		soundProcess.Process.Kill()
		soundProcess = nil
	}
}
