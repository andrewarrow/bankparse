package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
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
		HandleBothDays()
	} else if command == "move" {

		os.Rename("data/today.txt", "data/yesterday.txt")
		cmd := exec.Command("vim", "data/today.txt")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Run()
		HandleBothDays()

	} else if command == "help" {
		PrintHelp()
	}
}

func HandleBothDays() {
	yesterdayPending, yesterdayPosted := handleItems("data/yesterday.txt")
	todayPending, todayPosted := handleItems("data/today.txt")
	fmt.Println("<body>")
	fmt.Println("<table>")
	for k, v := range todayPending {
		if yesterdayPending[k] == nil {
			fmt.Printf("<tr><td>%s</td><td>%s</td><td></td></tr>\n", v.Thing, v.Amount)
		}
	}
	fmt.Println("</table>")
	fmt.Println("<table>")
	for k, v := range todayPosted {
		if yesterdayPosted[k] == nil {
			fmt.Printf("<tr><td>%s</td><td>%s</td><td></td></tr>\n", v.Thing, v.Amount)
		}
	}
	fmt.Println("</table>")
	fmt.Println("</body>")
}
