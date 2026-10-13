package main

import "sync"

type Book struct {
	ISBN      string `json:"isbn"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

type Library struct {
	mutex sync.RWMutex
	Books map[string]Book
}

func NewLibrary() *Library {
	return &Library{Books : CustomLibrary};
}

func (library *Library) GetAllBooks() []Book {
	library.mutex.RLock()
	defer library.mutex.RUnlock()
	var books []Book  
	for _, value := range library.Books {
		books = append(books, value)
	}
	return books
}

func (library *Library) GetBookByIsbn(isbn string) (Book, bool) {
	library.mutex.RLock()
	defer library.mutex.RUnlock()
	book, found := library.Books[isbn]
	return book, found
}

func (library *Library) AddBook(book Book) {
    library.mutex.Lock()
    defer library.mutex.Unlock()
    library.Books[book.ISBN] = book
}
