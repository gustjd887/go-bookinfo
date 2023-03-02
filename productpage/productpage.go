package main

import (
	"io/ioutil"
	"net/http"
)

type detail struct {
	Name      string
	Summary   string
	Type      string
	Page      int
	Publisher string
	Language  string
	Isbn10    string
	Isbn13    string
}

type review struct {
	Id       int
	Star     int
	Reviewer string
	Review   string
	Color    string
}

func main() {

	http.HandleFunc("/productpage", func(w http.ResponseWriter, r *http.Request) {
		detail := getJson("http://localhost:8002/detail")
		review := getJson("http://localhost:8001/review")
		w.Write(detail)
		w.Write(review)
	})
	http.ListenAndServe(":8003", nil)
}

func getJson(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return json
}
