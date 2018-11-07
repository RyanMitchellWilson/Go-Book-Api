package main

import (
	"time"
)

type Book struct {
	Author      string
	PublishDate time.Time
	Publisher   string
	Rating      int
	Status      string
	Title       string
}
