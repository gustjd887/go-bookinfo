package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type book struct {
	Id        int    `json:"Id"`
	Name      string `json:"Name"`
	Summary   string `json:"Summary"`
	Type      string `json:"Type"`
	Page      int    `json:"Page"`
	Publisher string `json:"Publisher"`
	Language  string `json:"Language"`
	Isbn10    string `json:"Isbn10"`
	Isbn13    string `json:"Isbn13"`
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "bookinfo"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/detail", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM detail")
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		// detail := []book{}
		detail := book{} // []type{}이 아니라 type{}으로 선언하여 json을 전달하려고 함.

		for rows.Next() {
			var r book
			err := rows.Scan(&r.Id, &r.Name, &r.Summary, &r.Type, &r.Page, &r.Publisher, &r.Language, &r.Isbn10, &r.Isbn13)
			if err != nil {
				panic(err)
			}
			// detail = append(detail, r)
			detail = r // rows가 한줄밖에 없기 떄문에 이렇게 처리. []json이 아니라 json type으로 전달.
		}

		bs, err := json.Marshal(detail)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(bs)
	})
	http.ListenAndServe(":8002", nil)
}
