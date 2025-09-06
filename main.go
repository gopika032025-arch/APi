package main

import (
	"ASS1GO/api"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	dsn := "root:Gopika@2001@tcp(127.0.0.1:3306)/testdb?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	fmt.Println("Connection established successfully!")

	api.RegisterRoutes(db)

	log.Println("Server started on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
