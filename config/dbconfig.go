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

		q := `CREATE TABLE Bookmarks
				(id INTEGER PRIMARY KEY AUTOINCREMENT, url TEXT, title TEXT, description TEXT, snapshot TEXT, date_added DATETIME, date_modified DATETIME, tags INTEGER);
			  CREATE TABLE Tags (id INTEGER PRIMARY KEY AUTOINCREMENT, bookmark_id INT NOT NULL, tag TEXT)
				`
		_, err = db.Exec(q)
		if err != nil {
			log.Fatal("Database: Error in setting up database tables.")
			return err
		}

		log.Print("Database: Database setup completed successfully")

	} else {
		log.Println("Database: bind.db found!")
	}

	return nil
}
