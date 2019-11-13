package main

import (
	"log"

	"github.com/facebookgo/inject"
)

type Book struct {
	title string
}

type BookService struct {
	database  Database  `inject:""`
	validator Validator `inject:""`
}

func (s *BookService) create(book Book) {
	log.Printf("creating new book entity: %v\n", book)
	if s.validator.validate(book) {
		s.database.save(book)
	}
}

type Database struct {
}

func (db *Database) save(entity interface{}) {
	log.Printf("saving entity to database: %v", entity)
}

type Validator struct {
}

func (v *Validator) validate(entity interface{}) bool {
	log.Printf("validating entity: %v", entity)
	return true
}

func main() {
	log.Println("starting injection example")
	var graph inject.Graph
	var service BookService
	graph.Provide(
		&inject.Object{Value: &service},
		&inject.Object{Value: Database{}},
		&inject.Object{Value: Validator{}},
	)
	graph.Populate()

	log.Println("listing known objects..")
	for _, obj := range graph.Objects() {
		log.Printf("object known: %v\n", obj)
	}

	service.create(Book{title: "Some book"})
}
