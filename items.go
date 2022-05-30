package main

import (
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

type Item struct {
	Thing  string
	Amount string
}

func handleItems(filename string) map[string]*Item {
	b, _ := ioutil.ReadFile(filename)
	s := string(b)
	tkn := html.NewTokenizer(strings.NewReader(s))
	tdCount := 0
	dateOn := false
	afterCount := 0
	thing := ""
	amount := ""
	items := map[string]*Item{}
	for {

		tt := tkn.Next()
		switch {

		case tt == html.ErrorToken:
			return items

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

				if afterCount == 0 {
					thing = txt
				} else {
					amount = txt
					parsedThing := ChargeParse(thing)
					items[parsedThing+"|"+amount] = &Item{parsedThing, amount}
				}
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

	return items
}
