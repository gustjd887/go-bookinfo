package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM rating")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	rating := []ratings{}

	for rows.Next() {
		var r ratings
		err := rows.Scan(&r.Id, &r.Star)
		if err != nil {
			panic(err)
		}
		rating = append(rating, r)
	}

	bs, err := json.Marshal(rating)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/rating", func(w http.ResponseWriter, r *http.Request) {
		w.Write(bs)
	})
	http.ListenAndServe(":8000", nil)
}
