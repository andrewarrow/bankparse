package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func handleItems(tkn *html.Tokenizer) {
	tdCount := 0
	dateOn := false
	afterCount := 0
	for {

		tt := tkn.Next()
		switch {

		case tt == html.ErrorToken:
			return

		case tt == html.StartTagToken:

			t := tkn.Token()
			if t.Data == "tr" {
				tdCount = 0
			} else if t.Data == "td" {
				tdCount++
			}

		case tt == html.TextToken:

			t := tkn.Token()
			txt := strings.TrimSpace(t.Data)
			if txt == "" {
				continue
			}
			if dateOn {
				if txt == "Print Details" || txt == "Posted" {
					dateOn = false
					afterCount = 0
					continue
				}
				fmt.Println(afterCount, txt)
				afterCount++
				if afterCount == 2 {
					dateOn = false
					afterCount = 0
				}
			}
			if len(txt) == 8 && string(txt[2]) == "/" && string(txt[5]) == "/" { // 05/19/22
				dateOn = true
				afterCount = 0
			}

		}

	}
}
