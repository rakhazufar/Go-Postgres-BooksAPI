package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-postgres/pkg/models"
	"github.com/rakhazufar/go-postgres/pkg/utils"
)

var NewBook models.Book

func GetBook (w http.ResponseWriter, r *http.Request) {
	newBooks, er := models.GetAllBooks()
	if er != nil {
		log.Fatalf("Error Get Books: %v", er)
	}
	res, err := json.Marshal(newBooks)
	if err != nil {
		log.Fatalf("Error when Marshal JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatalf("Error when parsing data")
	}

	bookDetails, _ := models.GetBookById(Id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
} 

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b:=CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatalf("Error when parsing data")
	}

	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)
    bookId, err := strconv.Atoi(params["bookId"])
	if err != nil {
		fmt.Fprintf(w, "Error: %s is not a valid book ID", params["bookId"])
		return
	}
	
	var NewBook models.Book

	if err := json.NewDecoder(r.Body).Decode(&NewBook); err != nil {
		fmt.Fprintf(w, "Error when decoding book data: %v", err)
		return
	}

	NewBook.ID = uint(bookId)
	
	updatedBook, err := models.UpdateBook(&NewBook)
	if err != nil {
		fmt.Printf("Error when decoding book data: %v", err)
	}

	res, _ := json.Marshal(updatedBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
