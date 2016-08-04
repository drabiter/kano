package main

import (
	"fmt"

	"github.com/apcera/termtables"
)

// TODO
const token string = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIzMGYxMjA0Mi1kNjIyLTQ2ZTgtYWQwYi00OWIzNGIwZWU4YWUiLCJzY29wZSI6WyJhbGwiXSwic3ViIjo1MTEwMCwiaXNzIjoxNDY4NDYzNDU1LCJleHAiOjE0NzY0MTIyNTV9.zIhaFLJXRnibGoTmnthAmGBjUM8jMhOAWJERNHSXCzI"
const username string = "akhemi"

func ListWatching() {
	records := ListAnime(username, token, STATUS_WATCHING)

	table := termtables.CreateTable()

	table.AddHeaders("ID", "Title", "Progress", "Total")
	for _, r := range records {
		table.AddRow(r.Anime.Slug, r.Anime.Title, r.EpisodesWatched, r.Anime.EpisodeCount)
	}

	fmt.Println(table.Render())
}
