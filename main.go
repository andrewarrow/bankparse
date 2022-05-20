package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  bank help         # this menu")
	fmt.Println("  bank today        # parse today")
	fmt.Println("  bank move         # move today")
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "today" {
		b, _ := ioutil.ReadFile("data/today.txt")
		s := string(b)
		tkn := html.NewTokenizer(strings.NewReader(s))
		handleItems(tkn)
	} else if command == "move" {
		os.Rename("data/today.txt", "data/yesterday.txt")
	} else if command == "help" {
		PrintHelp()
	}
}
