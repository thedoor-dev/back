package models

import "time"

type Post struct {
	ID         int64     `db:"pid"      json:"pid"`
	Title      string    `db:"title"    json:"title"`
	Abstract   string    `db:"abstract" json:"abstract,omitempty"`
	Article    string    `db:"article"  json:"article,omitempty"`
	Public     bool      `db:"public"   json:"public,omitempty"`
	CreateTime time.Time `db:"ctime"    json:"ctime"`
}

type PostList struct {
	Post
	Tags TagArr `json:"tags"`
}
