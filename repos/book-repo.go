package repos

import (
	"fmt"
	"go-crud/entities"
)

type BookRepo struct {
	books []entities.Book
}

func NewBookRepo() *BookRepo {
	return &BookRepo{make([]entities.Book, 0)}
}

func (b *BookRepo) Create(partial entities.Book) entities.Book {
	data := partial
	data.ID = uint(len(b.books)) + 1
	b.books = append(b.books, data)

	return data
}

func (b *BookRepo) GetOne(id uint) (entities.Book, error) {
	for _, book := range b.books {
		if book.ID == id {
			return book, nil
		}
	}

	return entities.Book{}, fmt.Errorf("book with id '%d' not found", id)
}
