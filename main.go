package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	c "github.com/GeistInDerSH/clearscreen"
	valid "github.com/asaskevich/govalidator"
	"github.com/fatih/color"
	"github.com/pelletier/go-toml"
	search "hashm.tech/moviefetch/bitsearchapi"
	"hashm.tech/moviefetch/chars"
	fetch "hashm.tech/moviefetch/fetchmovies"
	"hashm.tech/moviefetch/input"
	"hashm.tech/moviefetch/torrent"
	player "hashm.tech/moviefetch/watch"
)

var (
	config, _   = toml.LoadFile("config.toml")
	movies      = map[int]string{}
	torrents    = map[int]string{}
	quality     = config.Get("bitsearch.quality").(string)
	sort        = config.Get("bitsearch.sort").(string)
	apiKey      = config.Get("api.key").(string)
	searchlimit = config.Get("moviesearch.searchlimit").(int64)
	mediaplayer = config.Get("player.player").(string)
	lowseed     = config.Get("torrent.lowval").(int64)
	startscreen = config.Get("start.showscreen").(bool)
	bytestl     = config.Get("player.btl").(int64)
	g           = color.New(color.FgGreen)
	r           = color.New(color.FgRed)
)

func info() {
	fmt.Println("You can find this software for free @ github.com/MD5-Hashm/moviefetch")
	fmt.Println("I am NOT responsible for any legal issues or otherwise that you incurred using this software")
	os.Exit(0)
}

func lowval(v string) string {
	if strings.Contains(v, ".") || strings.Contains(strings.ToLower(v), "k") {
		return ""
	}
	val, _ := strconv.Atoi(v)
	if int64(val) < lowseed {
		return chars.Get("warning") + " "
	} else {
		return ""
	}
}

func getmovieinput(text string, limit int) int {
	for {
		inp := input.Get(text)
		if strings.ToLower(inp) == "s" {
			return -8
		}
		if valid.IsInt(inp) {
			iselection, _ := strconv.Atoi(inp)
			if iselection > limit {
				fmt.Println("Out of range :(")
			} else {
				return iselection
			}
		} else {
			fmt.Println("Not an int :(")
		}
	}
}

func searchfortorrent(movie string, mode string) {
	var titles []string
	var links []string
	var seeders []string
	var leechers []string
	var sizes []string
	if len(mode) <= 1 {
		titles, links, seeders, leechers, sizes = search.ParseWeb(search.GetWeb(search.MkUrl(movie+" "+quality, sort)))
	} else {
		titles, links, seeders, leechers, sizes = search.ParseWeb(search.GetWeb(search.MkUrl(movie, sort)))
	}
	fmt.Println("--------------------------------------------------------------------------------------------------------")
	for index := range titles {
		if strings.Contains(titles[index], "No results") || len(titles) < 10 {
			titles, links, seeders, leechers, sizes = search.ParseWeb(search.GetWeb(search.MkUrl(movie, sort)))
			if strings.Contains(titles[index], "No results") {
				fmt.Println(strings.Replace(titles[index], "+", " ", -1))
				os.Exit(1)
			}
		}
		fmt.Printf(strconv.Itoa(index+1) + ". " + titles[index] + " | ")
		g.Printf(seeders[index])
		fmt.Printf(" " + chars.Get("up") + " " + lowval(seeders[index]) + "| ")
		r.Printf(leechers[index])
		fmt.Printf(" " + chars.Get("down") + " | " + sizes[index] + "\n")
		torrents[index+1] = links[index]
		if index < 20 {
			fmt.Println("--------------------------------------------------------------------------------------------------------")
		}
	}
}

func editprefs() int {
	if strings.Contains(runtime.GOOS, "windows") {
		cmd := exec.Command("notepad.exe", "config.toml")
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	} else {
		cmd := exec.Command("open", "config.toml")
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
	return 0
}

func main() {
	var prequery string = ""
	var query string
	if len(os.Args) >= 2 && strings.Contains(os.Args[1], "--info") {
		info()
	} else if len(os.Args) >= 2 && !strings.Contains(os.Args[1], "--") {
		for index, arg := range os.Args[1:] {
			if index != len(os.Args)-1 {
				prequery += " " + arg
			} else {
				prequery += arg
			}
		}
	}
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		r.Printf("'config.toml' not found! Please create one (examples @ github)")
	}
	if prequery == "" {
		if startscreen {
			s := input.Get("enter/c: Continue\ne: Edit config\nex: Exit\n:")
			if strings.ToLower(s) == "c" || s == "" {
			} else if strings.ToLower(s) == "e" {
				_ = editprefs()

				quality, _ = config.Get("bitsearch.quality").(string)
				sort, _ = config.Get("bitsearch.sort").(string)
				apiKey, _ = config.Get("api.key").(string)
				searchlimit, _ = config.Get("moviesearch.searchlimit").(int64)
				mediaplayer, _ = config.Get("player.player").(string)
				lowseed, _ = config.Get("torrent.lowval").(int64)
				g.Println("Loaded changes")
				time.Sleep(time.Second)
			} else if strings.ToLower(s) == "ex" {
				os.Exit(0)
			} else {
				fmt.Println("Unrecognized command... continuing")
				time.Sleep(time.Second * 2)
			}
		}
		c.ClearScreen()
		color.Cyan("MovieFetch / MD5-Hashm <3")
		fmt.Println("-------------------------")
		query = input.Get("Movie Title: ")
	} else {
		query = prequery
	}
	c.ClearScreen()
	for num, movie := range fetch.FetchMovies(query, apiKey, searchlimit) {
		fmt.Printf(strings.Split(movie, ";")[0])
		movies[num+1] = strings.Split(movie, ";")[1]
	}
	fmt.Println("\ns: search with current query")
	iselection := getmovieinput(": ", len(movies))
	c.ClearScreen()
	if iselection == -8 {
		year := input.Get("Year? (optional) ")
		if year == "" {

		} else {
			query += " " + year
		}
		c.ClearScreen()
		searchfortorrent(query, "manual")
	} else {
		pmovie := fetch.FetchMovieById(movies[iselection], apiKey)
		fmt.Println(strings.Replace(pmovie, ";", "", -1))
		movie := strings.Split(pmovie, ";")[0]
		fmt.Println()
		fmt.Println("enter: search")
		//sselection := input.Get(": ")
		input.Get(": ")
		c.ClearScreen()
		/*if strings.ToLower(sselection) == "s" || strings.ToLower(sselection) == "" {
			searchfortorrent(strings.Replace(movie,":","",-1), "")
		}*/
		searchfortorrent(strings.Replace(movie, ":", "", -1), "")
	}
	torrentselection, terr := input.GetInt(": ", len(torrents))
	if terr != 0 {
		panic("Int was not passed and an unknown error occored")
	}
	c.ClearScreen()
	torrentf := torrents[torrentselection]
	tout := make(chan torrent.Stat)
	go torrent.AddTorrent(torrentf, tout)
	pl := false
	for {
		isMovie := true
		base := strings.Split(strings.Replace(strings.Replace(fmt.Sprintf("%v", <-tout), "{", "", -1), "}", "", -1), " ")
		if !pl {
			dlsize, _ := strconv.Atoi(strings.Split(fmt.Sprintf("%v", <-tout), " ")[1])
			if dlsize > int(bytestl) && isMovie {
				player.Launch(mediaplayer)
				pl = true
			}
		}
		if strings.Contains(fmt.Sprintf("%v", <-tout), "Stopped") {
			g.Println("Finished!")
			break
		}
		fmt.Printf("Status: %s\nTotal: %s\nDownloaded: %s\nPeers: %s\n", base[0], base[3], base[1], base[2])
		//fmt.Println(<-tout)
	}
	input.Get("Press enter to exit...")
}
