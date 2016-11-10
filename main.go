package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"errors"

	"github.com/chzyer/readline"
)

func main() {
	completer := readline.NewPrefixCompleter(
		readline.PcItem("list"),
		readline.PcItem("history"),
		readline.PcItem("bump"),
		readline.PcItem("search"),
		readline.PcItem("add"),
		readline.PcItem("delete"), // dup of remove
		readline.PcItem("finish"),
		readline.PcItem("rate"),
		readline.PcItem("quit"),
		readline.PcItem("exit"), // dup of quit
	)

	reader, err := readline.NewEx(&readline.Config{
		Prompt:       "> ",
		AutoComplete: completer,
	})

	if err != nil {
		panic(err)
	}
	defer reader.Close()

	if IsLoggedIn() {
		ListWatching()
	} else {
		fmt.Println("You're not logged in")

		Authenticate(loginParser())
		if IsLoggedIn() {
			ListWatching()
		}
	}

	for {
		line, err := reader.Readline()

		if err != nil {
			break
		}

		args := strings.Split(line, " ")
		if len(args) == 0 {
			continue
		}
		cmd := args[0]

		switch cmd {
		case "list":
			ListWatching()
		case "history", "record":
			ListCompleted()
		case "bump":
			id, count, err := bumpParser(args)

			if err != nil {
				fmt.Println(err)
				continue
			}

			UpdateEpisode(id, count)
		case "search":
			searchTerm := args[:]
			SearchTitle(strings.Join(searchTerm[1:], " "))
		case "add":
			id, err := strconv.ParseInt(args[1], 10, 0)

			if err != nil {
				fmt.Println("id must be number")
				continue
			}
			AddTitle(int(id))
		case "delete", "remove":
			id, err := strconv.ParseInt(args[1], 10, 0)

			if err != nil {
				fmt.Println("id must be number")
				continue
			}
			RemoveTitle(int(id))
		case "rate":
			id, err := strconv.ParseInt(args[1], 10, 0)

			if err != nil {
				fmt.Println("id must be number")
				continue
			}
			rating, err := strconv.ParseFloat(args[2], 32)
			if err != nil || !validRating(rating) {
				fmt.Println("rate must be 0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5")
				continue
			}
			RateTitle(int(id), float32(rating))
		case "finish":
			id, err := strconv.ParseInt(args[1], 10, 0)

			if err != nil {
				fmt.Println("id must be number")
				continue
			}
			FinishTitle(int(id))
		case "info":
			Info()
		case "quit", "exit":
			os.Exit(0)
		default:
		}
	}
}

func bumpParser(args []string) (int, int, error) {
	id, err := strconv.ParseInt(args[1], 10, 0)

	if err != nil {
		return -1, -1, errors.New("id must be number")
	}

	var count int64 = 1
	if len(args) >= 3 {
		count, err = strconv.ParseInt(args[2], 10, 0)

		if err != nil {
			return -1, -1, errors.New("number of episodes must be number (doh)")
		}
	}

	return int(id), int(count), nil
}

func loginParser() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	fmt.Println()

	return strings.TrimSpace(username), strings.TrimSpace(password)
}

func validRating(rating float64) bool {
	return math.Mod(rating, 0.5) == 0
}
