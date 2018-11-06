package main

import (
	"time"
)

type Book struct {
	Title       string
	Author      string
	Publisher   string
	PublishDate time.Time
	Rating      int
	Status      string
}
