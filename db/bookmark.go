package db

import "time"

type Bookmark struct {
	Id            uint      `json:"id"`
	Url           string    `json:"url"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Snapshot      string    `json:"snapshot"`
	Date_added    time.Time `json:"date_added"`
	Date_modified time.Time `json:"date_modified"`
	Tags          string    `json:"tags"`
}
