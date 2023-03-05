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

	for rows.Next() {
		var id int
		var star int
		err := rows.Scan(&id, &star)
		if err != nil {
			panic(err)
		}

		fmt.Println(id, star)
	}

	reviewer1 := ratings{
		Id:   1,
		Star: 5,
	}
	reviewer2 := ratings{
		Id:   2,
		Star: 4,
	}

	reviewer := []ratings{reviewer1, reviewer2}
	bs, err := json.Marshal(reviewer)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/rating", func(w http.ResponseWriter, r *http.Request) {
		w.Write(bs)
	})
	http.ListenAndServe(":8000", nil)

}
