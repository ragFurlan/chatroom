package entity

import "time"

type Message struct {
	User      string
	Message   string
	Timestamp time.Time
}
