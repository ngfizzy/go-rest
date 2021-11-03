package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)


// Book Struct (Model)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Get all books
func getBooks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode((books))
}

func createBook(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(int(time.Now().UnixMilli()))

	books = append(books, book)

	json.NewEncoder(rw).Encode((book))
}
func getBook(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(rw).Encode(item)
			return
		}
	}

	json.NewEncoder(rw).Encode(&Book{})
}
func updateBook(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(int(time.Now().UnixMicro()))
			books = append(books, book)

			json.NewEncoder(rw).Encode(book)
		}

	}

}
func deleteBook(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1])
			break
		}
	}

	json.NewEncoder(rw).Encode(books)
}

var books []Book

func main() {
	router := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "44323544", Title: "Book One", Author: &Author{ID: "1", FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "443235544", Title: "Book Two", Author: &Author{ID: "2", FirstName: "Jane", LastName: "Doe"}})
	books = append(books, Book{ID: "3", Isbn: "443285544", Title: "Book Three", Author: &Author{ID: "3", FirstName: "Zoe", LastName: "Bond"}})
	books = append(books, Book{ID: "4", Isbn: "443275544", Title: "Book Four", Author: &Author{ID: "4", FirstName: "Zamyla", LastName: "Bond"}})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8080", router))

}
