package book

import "fmt"

type fileRepository struct {
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (r *fileRepository) Create(book Book) (Book, error) {
	fmt.Println("Create to file")

	return book, nil
}

func (r *fileRepository) FindAll() ([]Book, error) {
	var books []Book

	fmt.Println("Find all from file")

	return books, nil
}

func (r *fileRepository) FindById(id int) (Book, error) {
	var book Book

	fmt.Println("Find by id from file")

	return book, nil
}

