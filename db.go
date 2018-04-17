package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InsertData(word string, example string) {
	db, err := sql.Open("postgres", "dbname=eng-words sslmode=disable")
	checkErr(err)

	// insert data
	stmt, err := db.Prepare("INSERT INTO wordlist(word,example) VALUES($1,$2) RETURNING uid")
	checkErr(err)

	_, err = stmt.Exec(word, example)
	checkErr(err)
}

func SelectData() ([]string, []string) {
	var words []string
	var examples []string

	db, err := sql.Open("postgres", "dbname=eng-words sslmode=disable")

	rows, err := db.Query("SELECT word,example FROM wordlist;")
	checkErr(err)

	for rows.Next() {
		var word string
		var example string
		err = rows.Scan(&word, &example)
		checkErr(err)
		words = append(words, word)
		examples = append(examples, example)
	}

	return words, examples
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
