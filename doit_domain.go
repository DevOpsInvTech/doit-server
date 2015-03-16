package main

import "time"

type Domain struct {
	ID        int    `sql:"not null;unique;AUTO_INCREMENT"`
	Name      string `sql:"type:varchar"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}