package db

type Config struct {
	Id    uint   `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// snap - 0, 1
// open_links - 0, 1
