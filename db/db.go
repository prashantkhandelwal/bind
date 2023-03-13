package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/prashantkhandelwal/bind/webext"
)

func db() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "bind.db")

	if err != nil {
		log.Fatal("ERROR:Database: Error in opening database.")
	}

	return db, nil
}

func GetSnap(id int64, url string) (bool, error) {
	db, _ := db()

	defer db.Close()

	snap, err := webext.Snap(url)
	if err != nil {
		log.Fatal("ERROR:Database:Save() - Error in getting the snap for the url")
	}

	_, err = db.Exec("UPDATE Bookmarks SET snapshot=? where id=?", snap, id)
	if err != nil {
		log.Fatalf("ERROR:Database:Update - Unable to update the snapshot of the bookmark url - %s - %s", url, err.Error())
		return false, err
	}

	return true, nil
}

func Save(b *Bookmark) (int64, error) {

	db, _ := db()

	defer db.Close()

	result, err := db.Exec("INSERT INTO Bookmarks VALUES(NULL,?,?,?,NULL,?,NULL,?,?);",
		b.Url,
		b.Title,
		b.Description,
		//snap,
		time.Now(),
		b.Tags,
		false)

	if err != nil {
		log.Fatalf("ERROR:Database:Save - Unable to save the bookmark - %s - %s", b.Url, err.Error())
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("ERROR:Database:Save - Unable to get the bookmark id - %s - %s", b.Url, err.Error())
		return 0, nil
	}

	return id, nil
}

/*
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
}*/
