package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type book struct {
	Name      string `json:"Name"`
	Summary   string `json:"Summary"`
	Type      string `json:"Type"`
	Page      int    `json:"Page"`
	Publisher string `json:"Publisher"`
	Language  string `json:"Language"`
	Isbn10    string `json:"Isbn10"`
	Isbn13    string `json:"Isbn13"`
}

func main() {
	book1 := book{
		Name:      "The Comedy of Errors",
		Summary:   "The Comedy of Errors is one of William Shakespeare's early plays. It is his shortest and one of his most farcical comedies, with a major part of the humour coming from slapstick and mistaken identity, in addition to puns and word play.",
		Type:      "paperback",
		Page:      200,
		Publisher: "PublisherA",
		Language:  "English",
		Isbn10:    "1234567890",
		Isbn13:    "123-1234567890",
	}

	bs, err := json.Marshal(book1)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/detail", func(w http.ResponseWriter, r *http.Request) {
		w.Write(bs)
	})
	http.ListenAndServe(":8002", nil)
}
