package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func handleItems(tkn *html.Tokenizer) {
	for {

		tt := tkn.Next()
		switch {

		case tt == html.ErrorToken:
			return

		case tt == html.StartTagToken:

			t := tkn.Token()
			fmt.Println(t, t.Data)

		case tt == html.TextToken:

			t := tkn.Token()
			fmt.Println("hi", t, t.Data)

		}

	}
}
