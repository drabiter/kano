package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
)

func main() {
	reader, err := readline.New("> ")

	if err != nil {
		panic(err)
	}
	defer reader.Close()

	ListWatching()

	for {
		line, err := reader.Readline()

		if err != nil {
			break
		}

		args := strings.Split(line, " ")
		cmd := args[0]

		if cmd == "bump" {
			id, err := strconv.ParseInt(args[1], 10, 0)

			if err != nil {
				fmt.Println("id must be number")
				continue
			}

			var count int64 = 1
			if len(args) >= 3 {
				count, err = strconv.ParseInt(args[2], 10, 0)

				if err != nil {
					fmt.Println("number of episodes must be number (doh)")
					continue
				}
			}

			UpdateEpisode(int(id), int(count))
		}

		if cmd == "list" {
			ListWatching()
		}
	}
}
