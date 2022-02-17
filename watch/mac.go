package watch

import (
	"os/exec"
)

const (
	qt = "open"
)

func MacWatchFile(filepath string) {
	cmd := exec.Command(qt, filepath)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
