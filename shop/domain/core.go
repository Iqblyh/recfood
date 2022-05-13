package domain

import "time"

type Shop struct {
	Id        int
	Shop_Name string
	Address   string
	Open      string
	Close     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
