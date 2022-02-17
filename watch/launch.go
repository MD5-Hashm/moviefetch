package watch

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"hashm.tech/moviefetch/input"
)

var (
	files = map[int]string{}
)

func ListCont(path string) []string {
	var dirs []string
	var files []string
	path += "/"
	fs, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range fs {
		if file.IsDir() {
			dirs = append(dirs, path+file.Name())
		} else {
			files = append(files, path+file.Name())
		}
	}
	for {
		var dirlen = len(dirs)
		var dirname string
		for _, dirname = range dirs {
			fs, err = ioutil.ReadDir(dirname)
			if err != nil {
				panic(err)
			}
		}
		for _, file := range fs {
			if file.IsDir() {
				dirs = append(dirs, dirname+"/"+file.Name())
			} else {
				files = append(files, dirname+"/"+file.Name())
			}
		}
		if len(dirs) > dirlen {
		} else {
			break
		}
	}
	return files
}

func Launch(player string) {
	for index, file := range ListCont("downloads") {
		fmt.Println(strconv.Itoa(index+1), file)
		files[index+1] = file
	}
	fs, _ := input.GetInt(": ", len(files))
	file := files[fs]
	Watch(player, file)
}

func Watch(player string, file string) {
	if player == "vlc" {
		VlcWatchFile(file)
	} else if player == "mpv" {
		MpvWatchFile(file)
	} else if player == "mac" || player == "quicktime" || player == "qt" {
		MacWatchFile(file)
	} else if player == "bundled" || player == "bundled-vlc" {
		BundledVlcWatchFile(file)
	}
}
