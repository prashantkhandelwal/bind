package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() error {
	if _, err := os.Stat("./bind.db"); err != nil {

		log.Println("Database: bind.db does not exist....Creating")

		db, err := sql.Open("sqlite3", "bind.db")
		if err != nil {
			log.Fatal("Error in opening/creating the database")
			return err
		}

		defer db.Close()

		log.Println("Database: bind.db created!")

		//q := `CREATE TABLE Bookmarks
		//		(id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, url TEXT NOT NULL, title TEXT NULL, description TEXT NULL, snapshot TEXT NULL, date_added DATETIME, date_modified DATETIME NULL, tags INTEGER NULL);
		q := `CREATE TABLE "config" (
					"id"	INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
					"key"	TEXT NOT NULL UNIQUE,
					"value"	TEXT NOT NULL,
					PRIMARY KEY("id" AUTOINCREMENT)
				);
				CREATE TABLE "bookmarks" (
					"id"	INTEGER NOT NULL UNIQUE,
					"url"	TEXT NOT NULL,
					"title"	BLOB NOT NULL,
					"description"	TEXT,
					"snapshot"	REAL,
					"date_added"	DATETIME NOT NULL,
					"date_modified"	DATETIME,
					"tags"	INTEGER,
					"is_archived"	INTEGER NOT NULL,
					PRIMARY KEY("id" AUTOINCREMENT)
				);`

		_, err = db.Exec(q)
		if err != nil {
			db.Close()
			err := os.Remove("bind.db")
			if err != nil {
				log.Fatalf("Setup failed - Cannot delete database %s", err)
			}
			log.Fatal("Database: Error in setting up database tables.")
			return err
		}

		log.Print("Database: Database setup completed successfully")

	} else {
		log.Println("Database: bind.db found!")
	}

	return nil
}
