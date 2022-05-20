package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
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
		yesterday := handleItems("data/yesterday.txt")
		today := handleItems("data/today.txt")
		fmt.Println("<body><table>")
		for k, v := range today {
			if yesterday[k] == nil {
				fmt.Printf("<tr><td>%s</td><td>%s</td><td></td></tr>\n", v.Thing, v.Amount)
			}
		}
		fmt.Println("</table></body>")
	} else if command == "move" {
		os.Rename("data/today.txt", "data/yesterday.txt")
	} else if command == "help" {
		PrintHelp()
	}
}
