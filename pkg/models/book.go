package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/rakhazufar/go-postgres/pkg/config"
	"gorm.io/gorm"
)


var db *gorm.DB


type Book struct {
	gorm.Model        
	Name string `json:"name"`
	Page int `json:"page"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatalf("error in miggration: %v", err)
	}
}

func (b *Book) CreateBook() *Book {
	result := db.Create(&b)
	if result.Error != nil {
		log.Fatalf("Error creating new Book: %s", result.Error)
	}
	return b 
}

func GetAllBooks()  ([]Book, error) {
	if db == nil {
		return nil, errors.New("database connection is not initialized")
	}

	var Books []Book
	result := db.Find(&Books)
	if result.Error != nil {
		// Mengembalikan error jika ada masalah saat mengambil data dari database
		return nil, result.Error
	}
	return Books, nil
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db:=db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) error {
	var book Book
	result := db.Where("ID=?", ID).Delete(&book)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}

func UpdateBook (book *Book) (*Book, error){
	result := db.Save(book)
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}