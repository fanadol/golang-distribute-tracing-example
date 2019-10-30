package server

import (
	"fmt"

	"github.com/fanadol/golang-distribute-tracing-example/models"
)

type Database struct {
	Post []models.Post
}

func NewDatabase() *Database {
	return &Database{Post: []models.Post{}}
}

func (d *Database) Create(title, body string) error {
	d.Post = append(d.Post, models.Post{Title: title, Body: body})
	fmt.Printf("%+v", d.Post)
	return nil
}
