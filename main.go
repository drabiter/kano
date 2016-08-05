package main

import (
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

		if args[0] == "bump" {
			count, _ := strconv.ParseInt(args[1], 10, 0)

			Update(int(count), 1)
		}
		if args[1] == "list" {
			ListWatching()
		}
	}
}
