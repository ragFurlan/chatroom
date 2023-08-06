package entity

import "time"

type Message struct {
	ID        int
	UserName  string
	Message   string
	Room      string
	Timestamp time.Time
}
