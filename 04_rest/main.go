package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)

type Librarian struct {
	library *Library
}

type librarianFunction func(http.ResponseWriter, *http.Request)

func (librarian *Librarian) GetBookByIsbn(writer http.ResponseWriter, request *http.Request) {
	var isbn string = request.URL.Query().Get("isbn")
	var book Book
	var found bool 
	book, found = librarian.library.GetBookByIsbn(isbn)
	if !found {
		http.NotFound(writer, request)
		return
	} 
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(book)
}

func (librarian *Librarian) PostBook(writer http.ResponseWriter, request *http.Request) {
	var newBook Book
	var err error = json.NewDecoder(request.Body).Decode(&newBook)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	librarian.library.AddBook(newBook)
	writer.WriteHeader(http.StatusCreated)
	fmt.Printf("%+v\n", newBook)
}

func (librarian *Librarian) HandleRequest(writer http.ResponseWriter, request *http.Request) {
	var methods map[string]librarianFunction = map[string]librarianFunction{
		http.MethodGet: librarian.GetBookByIsbn,
		http.MethodPost: librarian.PostBook,
	}
	var function librarianFunction
	var found bool
	function, found = methods[request.Method]
	if !found {
		http.Error(writer, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
	function(writer, request) 
}

func main() {
	var librarian *Librarian = &Librarian{library : NewLibrary()}
	http.HandleFunc("/bookISBN", librarian.HandleRequest)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}