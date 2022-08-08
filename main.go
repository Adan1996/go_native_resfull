package main

import (
	"encoding/json"
	"net/http"
)

type Books struct {
	Title string `json:"name"`
	ID    string `json:"id"`
	Price int    `json:"price"`
	Qty   int    `json:"qty"`
}

type bookHandler struct {
	store map[string]Books
}

func (b *bookHandler) get(w http.ResponseWriter, r *http.Request) {

	books := make([]Books, len(b.store))

	i := 0
	for _, book := range b.store {
		books[i] = book
		i++
	}

	jsonBytes, err := json.Marshal(books)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func newBooksHandler() *bookHandler {
	return &bookHandler{
		store: map[string]Books{
			"id1": Books{
				Title: "Jago Masak Dalam Sehari",
				ID:    "id1",
				Price: 129000,
				Qty:   20,
			},
		},
	}
}

func main() {
	booksHandler := newBooksHandler()
	http.HandleFunc("/books", booksHandler.get)
	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}
