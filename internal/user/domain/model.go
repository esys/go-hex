package domain

import "time"

type User struct {
	ID      string     `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}
