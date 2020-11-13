package models
//
//import (
//	"github.com/GoAdminGroup/go-admin/modules/db"
//	DB "github.com/natalizhy/pupils_grade_sheet/db"
//
//	"github.com/jinzhu/gorm"
//)
//
//
//type Students struct {
//	//Id          string `json:"id"`
//	Name        string `gorm:""json:"name"`
//	Author      string `json:"author"`
//	Publication string `json:"publication"`
//}
//
//func init() {
//	config.Connect()
//	db = config.GetDB()
//	db.AutoMigrate(&Book{})
//}
//
//func (s *Students) CreateSchools() *Students {
//	DB.NewRecord(b)
//	db.Create(&b)
//	return b
//}
//
//func  GetAllBooks() []Book {
//	var Books []Book
//	db.Find(&Books)
//	return Books
//}
//
//func GetBookById(Id int64) (*Book , *gorm.DB){
//	var getBook Book
//	db:=db.Where("ID = ?", Id).Find(&getBook)
//	return &getBook, db
//}
//
//func DeleteBook(ID int64) Book {
//	var book Book
//	db.Where("ID = ?", ID).Delete(book)
//	return book
//}