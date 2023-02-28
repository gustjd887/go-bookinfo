package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type reviewes struct {
	Id       int    `json:"Id"`
	Star     int    `json:"Star"`
	Reviewer string `json:"Reviewer"`
	Review   string `json:"Review"`
	Color    string `json:"color"`
}

func main() {
	reviewer1 := reviewes{
		Id:       1,
		Star:     0,
		Reviewer: "Reviewer1",
		Review:   "An extremely entertaining play by Shakespeare. The slapstick humour is refreshing!",
		Color:    "",
	}
	reviewer2 := reviewes{
		Id:       2,
		Star:     0,
		Reviewer: "Reviewer2",
		Review:   "Absolutely fun and entertaining. The play lacks thematic depth when compared to other plays by Shakespeare.",
		Color:    "",
	}

	reviewer := []reviewes{reviewer1, reviewer2}
	bs, err := json.Marshal(reviewer)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/review", func(w http.ResponseWriter, r *http.Request) {
		w.Write(bs)
	})
	http.ListenAndServe(":8001", nil)

}
