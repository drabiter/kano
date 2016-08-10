package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/apcera/termtables"
	"github.com/drabiter/kano/hummingbird"
	"github.com/drabiter/kano/utils"
	homedir "github.com/mitchellh/go-homedir"
)

type Config struct {
	User  string `json:"user"`
	Token string `json:"token"`
}

const kanoUserEnv string = "KANO_USER"
const kanoTokenEnv string = "KANO_TOKEN"
const kanoConfLoc string = "/.kano"

var cfg Config

var records []hummingbird.Record

func init() {
	readConfig(&cfg)
}

func IsLoggedIn() bool {
	return len(strings.TrimSpace(cfg.User)) > 0 && len(strings.TrimSpace(cfg.Token)) > 0
}

func Info() {
	fmt.Println(fmt.Sprintf("user  : %s", cfg.User))
	fmt.Println(fmt.Sprintf("Token : %s", cfg.Token))
}

func Authenticate(username string, password string) {
	auth, err := hummingbird.Authenticate(username, password)

	if err != nil {
		fmt.Println(err)
		return
	}

	cfg.Token = auth
	cfg.User = username

	writeConfig(&cfg)

	fmt.Println("Welcome!")
	fmt.Println()
}

func ListWatching() {
	records = hummingbird.ListAnimeWatching(cfg.User, cfg.Token)

	table := termtables.CreateTable()

	table.AddHeaders("ID", "Title", "Type", "Progress", "Total")
	for idx, r := range records {
		table.AddRow(idx, r.Anime.Title, r.Anime.ShowType, r.EpisodesWatched, r.Anime.EpisodeCount)
	}

	fmt.Println(table.Render())
}

func ListCompleted() {
	records = hummingbird.ListAnimeCompleted(cfg.User, cfg.Token)
	sort.Sort(utils.RecordLastWatchedDescSorter(records))

	table := termtables.CreateTable()

	table.AddHeaders("ID", "Title", "Type", "Total", "Rating", "Community Rating")
	for idx, r := range records {
		table.AddRow(idx, r.Anime.Title, r.Anime.ShowType, r.Anime.EpisodeCount, r.Rating.Value, r.Anime.Rating)
	}

	fmt.Println(table.Render())
}

func UpdateEpisode(index int, count int) {
	record := records[index]

	totalWatched := record.EpisodesWatched + count
	status := hummingbird.StatusWatching
	if totalWatched >= record.Anime.EpisodeCount {
		status = hummingbird.StatusCompleted
	}

	hummingbird.UpdateAnime(cfg.Token, record.Anime.ID, totalWatched, status)

	ListWatching()
}

func SearchTitle(keyword string) {
	result := hummingbird.SearchAnime(keyword)

	table := termtables.CreateTable()

	table.AddHeaders("ID", "Title", "Type", "Total", "Status", "Rating")
	for _, a := range result {
		table.AddRow(a.ID, a.Title, a.ShowType, a.EpisodeCount, a.Status, a.Rating)
	}

	fmt.Println(table.Render())
}

func AddTitle(id int) {
	status := hummingbird.StatusWatching
	totalWatched := 0

	hummingbird.UpdateAnime(cfg.Token, id, totalWatched, status)

	ListWatching()
}

func RemoveTitle(id int) {
	hummingbird.RemoveAnime(cfg.Token, id)

	ListWatching()
}

func FinishTitle(index int) {
	record := records[index]

	count := record.Anime.EpisodeCount - record.EpisodesWatched

	UpdateEpisode(index, count)
}

func RateTitle(index int, rating float32) {
	record := records[index]

	hummingbird.RateAnime(cfg.Token, record.Anime.ID, rating)
}

func try(err error) {
	if err != nil {
		panic(err)
	}
}

func configPath() string {
	dir, err := homedir.Dir()
	try(err)
	return dir + kanoConfLoc
}

func readConfig(cfg *Config) {
	buffer, err := ioutil.ReadFile(configPath())
	if err == nil {
		json.Unmarshal(buffer, cfg)
	} else {
		cfg = &Config{}
	}
}

func writeConfig(cfg *Config) {
	buffer, err := json.Marshal(cfg)
	try(err)

	err = ioutil.WriteFile(configPath(), buffer, 0644)
	try(err)
}
