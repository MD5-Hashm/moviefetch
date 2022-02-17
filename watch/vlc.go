package watch

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

const (
	windowsPath = `C:\Program Files\VideoLAN\VLC\vlc.exe`
	macvlc      = `/Applications/VLC.app/Contents/MacOS/VLC`
	unixvlc     = `/usr/bin/vlc`
)

func VlcWatchFile(filepath string) {
	if strings.Contains(runtime.GOOS, "windows") {
		cmd := exec.Command(windowsPath, `"file:///`+filepath+`"`)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		color.Green("Launched vlc!")
	} else if strings.Contains(runtime.GOOS, "darwin") {
		cmd := exec.Command(macvlc, `"file:///`+filepath+`"`)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		color.Green("Launched vlc!")
	} else {
		cmd := exec.Command(unixvlc, `"file:///`+filepath+`"`)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		color.Green("Launched vlc!")
	}
}

func BundledVlcWatchFile(filepath string) {
	cmd := exec.Command(`VLC\vlc.exe`, `"file:///`+filepath+`"`)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	color.Green("Launched vlc!")
}
