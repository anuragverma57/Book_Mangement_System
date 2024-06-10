package models

import (
	"github.com/anuragverma57/Book_Mangement_System/pkg/config"
	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	// db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var Book Book
	db := db.Where("Id=?", Id).Find(&Book)
	return &Book, db
}

func DeleteBook(Id int64) Book {
	var Book Book
	db.Where("Id=?", Id).Delete(Book)
	return Book
}
