package server

import (
	"context"

	"github.com/fanadol/golang-distribute-tracing-example/models"
	"github.com/opentracing/opentracing-go"
)

type Database struct {
	Post []models.Post
}

func NewDatabase() *Database {
	return &Database{Post: []models.Post{}}
}

func (d *Database) Create(ctx context.Context, title, body string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "Server-Create-Database")
	defer span.Finish()

	d.Post = append(d.Post, models.Post{Title: title, Body: body})
	return nil
}

func (d *Database) Get(ctx context.Context) ([]models.Post, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Server-Get-Database")
	defer span.Finish()

	return d.Post, nil
}
