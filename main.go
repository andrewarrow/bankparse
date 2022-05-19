package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  bank help         # this menu")
	fmt.Println("  bank today        # parse today")
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
		lines := strings.Split(s, "\n")
		for _, line := range lines {
			tokens := strings.Split(line, ",")
			fmt.Println(tokens)
		}
	} else if command == "phases" {
	} else if command == "help" {
		PrintHelp()
	}
}
