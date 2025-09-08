package api

import (
	"ASS1GO/dataservice"
	"ASS1GO/model"
	"database/sql"
)

type IBizLogic interface {
	CreateBookLogic(book model.Book) error
}

type Bizlogic struct {
	DB *sql.DB
}

func NewBizLogic(db *sql.DB) *Bizlogic {
	return &Bizlogic{DB: db}
}

func (bl *Bizlogic) CreateBookLogic(book model.Book) error {
	//validations for get method
	if err := dataservice.CreateBook(bl.DB, book); err != nil {
		return err
	}
	return nil
}
