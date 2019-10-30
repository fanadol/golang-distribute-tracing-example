package server

import (
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
	return nil
}

func (d *Database) Get() ([]models.Post, error) {
	return d.Post, nil
}
