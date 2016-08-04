package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const STATUS_WATCHING string = "currently-watching"
const STATUS_COMPLETED string = "completed"
const STATUS_PLAN_TO_WATCH string = "plan-to-watch"
const STATUS_DROPPED string = "dropped"
const STATUS_HOLD string = "on-hold"

type Record struct {
	ID              int       `json:"id"`
	EpisodesWatched int       `json:"episodes_watched"`
	LastWatched     time.Time `json:"last_watched"`
	Status          string    `json:"status"`
	Anime           struct {
		ID           int    `json:"id"`
		Slug         string `json:"slug"`
		Status       string `json:"status"`
		Title        string `json:"title"`
		EpisodeCount int    `json:"episode_count"`
	} `json:"anime"`
	Rating struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"rating"`
}

func ListAnime(username string, token string, status string) []Record {
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://hummingbird.me/api/v1/users/%s/library", username), nil)

	query := req.URL.Query()
	query.Add("auth_token", token)
	query.Add("status", status)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}

	res, _ := client.Do(req)
	defer res.Body.Close()

	var records []Record
	bytes, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bytes, &records)

	return records
}
