package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Book struct {
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books []Book

func main() {
    books = []Book{
        {Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
        {Title: "To Kill a Mockingbird", Author: "Harper Lee"},
    }

    http.HandleFunc("/books", handleBooks)

    fmt.Println("Server starting on port 8000...")
    http.ListenAndServe(":8000", nil)
}

func handleBooks(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        json.NewEncoder(w).Encode(books)
    case "POST":
        var book Book
        json.NewDecoder(r.Body).Decode(&book)
        books = append(books, book)
        json.NewEncoder(w).Encode(book)
    }
}
