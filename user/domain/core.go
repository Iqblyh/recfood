package domain

import "time"

type User struct {
	Username  string
	Email     string
	Password  string
	Fullname  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
