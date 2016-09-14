package hummingbird

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const apiUrl string = "https://hummingbird.me/api/v1"
const StatusWatching string = "currently-watching"
const StatusCompleted string = "completed"
const StatusPlanToWatch string = "plan-to-watch"
const StatusDropped string = "dropped"
const StatusHold string = "on-hold"

type Anime struct {
	ID           int     `json:"id"`
	Slug         string  `json:"slug"`
	Status       string  `json:"status"`
	Title        string  `json:"title"`
	EpisodeCount int     `json:"episode_count"`
	ShowType     string  `json:"show_type"`
	Rating       float64 `json:"community_rating"`
}

type Record struct {
	ID              int       `json:"id"`
	EpisodesWatched int       `json:"episodes_watched"`
	LastWatched     time.Time `json:"last_watched"`
	Status          string    `json:"status"`
	Anime           Anime     `json:"anime"`
	Rating          struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"rating"`
}

type Payload struct {
	AuthToken       string   `json:"auth_token"`
	Status          *string  `json:"status,omitempty"`
	EpisodesWatched *int     `json:"episodes_watched,omitempty"`
	SaneRating      *float32 `json:"rating,omitempty"`
}

type Credential struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func Authenticate(username string, password string) (string, error) {
	credential := &Credential{
		UserName: username,
		Password: password,
	}

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(credential)

	res, err := http.Post(fmt.Sprintf("%s/users/authenticate", apiUrl), "application/json; charset=utf-8", buffer)
	defer res.Body.Close()

	if err != nil {
		return "", err
	} else if res.StatusCode != 201 {
		return "", errors.New("Invalid credentials")
	}

	body, err := ioutil.ReadAll(res.Body)

	return strings.Replace(string(body), "\"", "", 2), nil
}

func ListAnimeCompleted(username string, token string) []Record {
	return ListAnime(username, token, StatusCompleted)
}

func ListAnimeWatching(username string, token string) []Record {
	return ListAnime(username, token, StatusWatching)
}

func ListAnime(username string, token string, status string) []Record {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/users/%s/library", apiUrl, username), nil)

	query := req.URL.Query()
	query.Add("auth_token", token)
	query.Add("status", status)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}

	res, _ := client.Do(req)
	defer res.Body.Close()

	var records []Record
	json.NewDecoder(res.Body).Decode(&records)

	return records
}

func UpdateAnime(token string, id int, episode int, status string) Record {
	return putLibraries(fmt.Sprintf("%s/libraries/%s", apiUrl, strconv.Itoa(id)), &Payload{
		AuthToken:       token,
		Status:          &status,
		EpisodesWatched: &episode,
	})
}

func RateAnime(token string, id int, rating float32) Record {
	return putLibraries(fmt.Sprintf("%s/libraries/%s", apiUrl, strconv.Itoa(id)), &Payload{
		AuthToken:  token,
		SaneRating: &rating,
	})
}

func SearchAnime(keyword string) []Anime {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/search/anime", apiUrl), nil)

	query := req.URL.Query()
	query.Add("query", keyword)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	res, _ := client.Do(req)
	defer res.Body.Close()

	var animes []Anime
	json.NewDecoder(res.Body).Decode(&animes)

	return animes
}

func RemoveAnime(token string, id int) {
	payload := &Payload{
		AuthToken: token,
	}

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(payload)

	res, err := http.Post(fmt.Sprintf("%s/libraries/%s/remove", apiUrl, strconv.Itoa(id)), "application/json; charset=utf-8", buffer)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
}

func putLibraries(url string, payload *Payload) Record {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(payload)

	res, err := http.Post(url, "application/json; charset=utf-8", buffer)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	var record Record
	json.NewDecoder(res.Body).Decode(&record)

	return record
}
