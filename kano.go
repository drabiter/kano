package main

import (
	"fmt"

	"github.com/apcera/termtables"
)

// TODO
const token string = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIzMGYxMjA0Mi1kNjIyLTQ2ZTgtYWQwYi00OWIzNGIwZWU4YWUiLCJzY29wZSI6WyJhbGwiXSwic3ViIjo1MTEwMCwiaXNzIjoxNDY4NDYzNDU1LCJleHAiOjE0NzY0MTIyNTV9.zIhaFLJXRnibGoTmnthAmGBjUM8jMhOAWJERNHSXCzI"
const username string = "akhemi"

var records []Record

func ListWatching() {
	records = ListAnime(username, token, StatusWatching)

	table := termtables.CreateTable()

	table.AddHeaders("ID", "Title", "Progress", "Total")
	for idx, r := range records {
		table.AddRow(idx, r.Anime.Title, r.EpisodesWatched, r.Anime.EpisodeCount)
	}

	fmt.Println(table.Render())
}

func GetTitle(index int64) string {
	return records[index].Anime.Title
}

func Update(index int, count int) {
	record := records[index]
	UpdateAnime(token, record.Anime.ID, record.EpisodesWatched+count)

	ListWatching()
}
