package bitsearch

import (
	"io"
	"net/http"
	"strings"
)

const (
	baseUrl = "https://bitsearch.to/search?q="
)

func MkUrl(query string, sort string) string {
	query = strings.Replace(query, " ", "+", -1)
	if sort == "" {
		return baseUrl + query
	} else {
		return baseUrl + query + "&sort=" + sort
	}
}

func GetWeb(url string) (html string, u string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	out := new(strings.Builder)
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
	return out.String(), url
}

func ParseWeb(web string, url string) (title []string, magnet []string, seed []string, leech []string, siz []string) {
	// -13 is size, -11 is seeder, -8 is leecher
	var torrents []string
	var titles []string
	var seeders []string
	var leechers []string
	var size []string
	var temp []string
	var splitweb []string = strings.Split(web, "\n")
	for index, line := range splitweb {
		if strings.HasPrefix(line, `<a href="https://itorrents.org/torrent/`) {
			temp = nil
			torrents = append(torrents, strings.Split(splitweb[index], `"`)[1])
			temp = append(strings.Split(line, `[Bitsearch.to]`)[0:])
			titles = append(titles, strings.Split(temp[1], `data`)[0])

			splitweb[index-11] = strings.Replace(splitweb[index-11], `<font color="#0AB49A">`, "", -1)
			splitweb[index-11] = strings.Replace(splitweb[index-11], `</font>`, "", -1)
			splitweb[index-11] = strings.Replace(splitweb[index-11], ` `, "", -1)
			seeders = append(seeders, splitweb[index-11])

			splitweb[index-8] = strings.Replace(splitweb[index-8], `<font color="#C35257">`, "", -1)
			splitweb[index-8] = strings.Replace(splitweb[index-8], `</font>`, "", -1)
			splitweb[index-8] = strings.Replace(splitweb[index-8], ` `, "", -1)
			leechers = append(leechers, splitweb[index-8])

			splitweb[index-13] = strings.Replace(splitweb[index-13], `<div><img src="/icons/size.svg" alt="Size" style="transform: translateY(-2px)">`, "", -1)
			splitweb[index-13] = strings.Replace(splitweb[index-13], `</div>`, "", -1)
			splitweb[index-13] = strings.Replace(splitweb[index-13], ` `, "", -1)
			size = append(size, splitweb[index-13])
		}
	}
	for num, title := range titles {
		titles[num] = strings.Replace(title, `"`, "", -1)
	}
	if len(torrents) < 1 {
		titles = append(titles, "No results found for "+strings.Split(strings.Split(url, "q=")[1], "&sort")[0]+" :( try searching manually", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "")
		torrents = append(torrents, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "")
		seeders = append(seeders, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "")
		leechers = append(leechers, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "")
		size = append(size, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "")
		return titles, torrents, seeders, leechers, size
	}
	return titles, torrents, seeders, leechers, size
}
