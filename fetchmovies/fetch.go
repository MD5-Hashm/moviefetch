package fetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func baseUrl(apiKey string) string {
	return "https://api.themoviedb.org/3/search/movie?api_key=" + apiKey + "&query="
}

/*var (
	genre = map[int]string{
		28:    "Action",
		12:    "Adventure",
		16:    "Animation",
		35:    "Comedy",
		80:    "Crime",
		99:    "Documentary",
		18:    "Drama",
		10751: "Family",
		14:    "Fantasy",
		36:    "History",
		27:    "Horror",
		10402: "Music",
		9648:  "Mystery",
		10749: "Romance",
		878:   "Science Fiction",
		10770: "TV Movie",
		53:    "Thriller",
		10752: "War",
		37:    "Western"}
)*/

type response struct {
	Page          int      `json:"page"`
	Results       []movies `json:"results"`
	Total_pages   int      `json:"total_pages"`
	Total_results int      `json:"total_results"`
}

type movies struct {
	Poster_path       string  `json:"poster_path"`
	Adult             bool    `json:"adult"`
	Overview          string  `json:"overview"`
	Release_date      string  `json:"release_date"`
	Genre_ids         []int   `json:"genre_ids"`
	Id                int     `json:"id"`
	Original_title    string  `json:"original_title"`
	Original_language string  `json:"original_language"`
	Title             string  `json:"title"`
	Backdrop_path     string  `json:"backdrop_path"`
	Popularity        float64 `json:"popularity"`
	Vote_count        int     `json:"vote_count"`
	Video             bool    `json:"video"`
	Vote_average      float64 `json:"vote_average"`
}

func FetchMovies(name string, apikey string, searchlimit int64) []string {
	name = strings.Replace(name, " ", "+", -1)
	url := baseUrl(apikey) + name

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response response
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	var res []string
	var count int64
	for num, movie := range response.Results {
		if count >= searchlimit {
			break
		}
		if len(movie.Release_date) < 4 {
			res = append(res, strconv.Itoa(num+1)+". "+movie.Title+"\n"+";"+strconv.Itoa(movie.Id))
		} else {
			res = append(res, strconv.Itoa(num+1)+". "+movie.Title+": "+movie.Release_date[0:4]+"\n"+";"+strconv.Itoa(movie.Id))
		}
		count++
	}

	if len(res) < 1 {
		res = append(res, "No results found;")
		return res
	}
	return res
}

func FetchMovieById(id string, apiKey string) string {
	url := "https://api.themoviedb.org/3/movie/" + id + "?api_key=" + apiKey
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var movie movies
	err = json.Unmarshal(body, &movie)
	if err != nil {
		panic(err)
	}
	var out string
	out += fmt.Sprintf("%s %s;\n\n%s\n\nVote Average: %.1f", movie.Title, movie.Release_date[0:4], movie.Overview, movie.Vote_average)
	return out
}
