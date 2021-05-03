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
	fmt.Println("  bank ls           # list servers")
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "ls" {
		b, _ := ioutil.ReadFile("TIAA_CHECKING8361_transactions.csv")
		s := string(b)
		lines := strings.Split(s, "\n")
		for _, line := range lines {
			tokens := strings.Split(line, ",")
			if len(tokens) <= 1 {
				continue
			}
			fmt.Println(tokens[0], tokens[1])
		}
	} else if command == "phases" {
	} else if command == "help" {
		PrintHelp()
	}
}
