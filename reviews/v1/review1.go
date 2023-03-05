package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "bookinfo"
)

type ratings struct {
	Id   int `json:"Id"`
	Star int `json:"Star"`
}

type reviewes struct {
	Id       int    `json:"Id"`
	Star     int    `json:"Star"`
	Reviewer string `json:"Reviewer"`
	Review   string `json:"Review"`
	Color    string `json:"color"`
}

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/review", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8000/rating")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var rat []ratings
		json.Unmarshal(data, &rat)

		rows, err := db.Query("SELECT * FROM review")
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		review := []reviewes{}

		for rows.Next() {
			var r reviewes
			err := rows.Scan(&r.Id, &r.Reviewer, &r.Review, &r.Color)
			if err != nil {
				panic(err)
			}

			review = append(review, r)
		}

		for i, v := range review {
			if v.Id == 1 {
				for _, v2 := range rat {
					if v2.Id == 1 {
						review[i].Star = v2.Star
						// review[i].Color = ""
					}
				}
			}
			if v.Id == 2 {
				for _, v2 := range rat {
					if v2.Id == 2 {
						review[i].Star = v2.Star
						// review[i].Color = ""
					}
				}
			}
		}

		bs, err := json.Marshal(review)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(bs)
	})
	http.ListenAndServe(":8001", nil)

}
