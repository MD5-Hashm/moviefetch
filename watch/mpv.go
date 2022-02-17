package watch

import (
	"os/exec"

	"github.com/fatih/color"
)

const (
	mpv = "mpv"
)

func MpvWatchFile(filepath string) {
	cmd := exec.Command(mpv, filepath)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	color.Green("Launched " + mpv + "!")
}
