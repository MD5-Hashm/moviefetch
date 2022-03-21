package watch

import (
	"os/exec"

	"github.com/fatih/color"
)

func WinDef(filepath string) {
	cmd := exec.Command(`"` + filepath + `"`)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	color.Green("Launched vlc!")
}
