package model

import "time"

type Book struct {
	id        int
	name      string
	author    string
	published time.Time
} 