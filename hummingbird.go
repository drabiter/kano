package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

const apiUrl string = "https://hummingbird.me/api/v1"
const StatusWatching string = "currently-watching"
const StatusCompleted string = "completed"
const StatusPlanToWatch string = "plan-to-watch"
const StatusDropped string = "dropped"
const StatusHold string = "on-hold"

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

type Payload struct {
	AuthToken       string `json:"auth_token"`
	Status          string `json:"status"`
	EpisodesWatched int    `json:"episodes_watched"`
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

func UpdateAnime(token string, id int, episode int) Record {
	payload := Payload{
		AuthToken:       token,
		Status:          StatusWatching,
		EpisodesWatched: episode,
	}

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(payload)
	// bytess, _ := json.Marshal(payload)

	fmt.Println(fmt.Sprintf("%s/libraries/%s", apiUrl, strconv.Itoa(id)))
	fmt.Printf("%+v\n", payload)

	res, err := http.Post(fmt.Sprintf("%s/libraries/%s", apiUrl, strconv.Itoa(id)), "application/json; charset=utf-8", buffer)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(os.Stdout, res.Body)
	defer res.Body.Close()

	var record Record
	json.NewDecoder(res.Body).Decode(&record)
	fmt.Printf("%+v\n", record)

	return record
}
