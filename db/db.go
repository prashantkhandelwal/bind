package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Bookmark struct {
	id            uint
	url           string
	title         string
	description   string
	snapshot      string
	date_added    time.Time
	date_modified time.Time
	tags          string
}

func AddTag() (bool, error) {
	db, err := sql.Open("sqlite3", "bind.db")
	if err != nil {
		log.Fatal("Database: Error in opening database.")
		return false, err
	}

	//statement, err := db.Prepare("insert into Bookmarks(url, title, description, snapshot, date_added, date_modified, tags)")
	statement, err := db.Prepare("insert into Tags(bookmark_id, tag) values (?,?)")
	if err != nil {
		log.Fatal("Database: Error in opening database.")
		return false, err
	}
	_, err = statement.Exec(1, "Github")

	if err != nil {
		log.Fatal("Database: Error in saving Tags.")
		return false, err
	}
	return true, nil
}
