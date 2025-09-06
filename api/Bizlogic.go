package api

import (
	"ASS1GO/dataservice"
	"database/sql"
	"net/http"
)

func CreateBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.CreateBook(db, w, r)
}
