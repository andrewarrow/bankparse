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

func handleItems(filename string) (map[string]*Item, map[string]*Item) {
	b, _ := ioutil.ReadFile(filename)
	s := string(b)
	tkn := html.NewTokenizer(strings.NewReader(s))
	tdCount := 0
	dateOn := false
	afterCount := 0
	thing := ""
	amount := ""
	postedOn := false

	pendingItems := map[string]*Item{}
	postedItems := map[string]*Item{}

	for {

		tt := tkn.Next()
		switch {

		case tt == html.ErrorToken:
			return pendingItems, postedItems

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
			if txt == "Posted" {
				postedOn = true
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
					if postedOn {
						postedItems[thing+"|"+amount] = &Item{thing, amount}
					} else {
						parsedThing := ChargeParse(thing)
						pendingItems[parsedThing+"|"+amount] = &Item{thing, amount}
					}
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

	return pendingItems, postedItems
}
