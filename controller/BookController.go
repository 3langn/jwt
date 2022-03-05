package controller

import "github.com/3langn/learn-go/models"

type IBookController interface {
	GetBook(id int) (book models.Book, err error)
	GetBooks() (books []models.Book, err error)
	AddBook(book models.Book) (err error)
	UpdateBook(book models.Book) (err error)
	DeleteBook(id int) (err error)
}


type BookController struct {
}

func (bookCtrl BookController) GetBook(id int) (book models.Book, err error) {
	bookSnapshot := models.GetDb().Where("id = ?", id).First(&book)
	if bookSnapshot.Error != nil {
		return book, bookSnapshot.Error
	}
	return book, nil
}

func (bookCtrl BookController) GetBooks() (books []models.Book, err error) {
	bookSnapshot := models.GetDb().Find(&books)
	if bookSnapshot.Error != nil {
		return books, bookSnapshot.Error
	}
	return books, nil
}

func (bookCtrl BookController) AddBook(book models.Book) (err error) {
	bookSnapshot := models.GetDb().Create(&book)
	if bookSnapshot.Error != nil {
		return bookSnapshot.Error
	}
	return nil
}

func (bookCtrl BookController) UpdateBook(book models.Book) (err error) {
	bookSnapshot := models.GetDb().Save(&book)
	if bookSnapshot.Error != nil {
		return bookSnapshot.Error
	}
	return nil
}

func (bookCtrl BookController) DeleteBook(id int) (err error) {
	bookSnapshot := models.GetDb().Where("id = ?", id).Delete(&models.Book{})
	if bookSnapshot.Error != nil {
		return bookSnapshot.Error
	}
	return nil
}
