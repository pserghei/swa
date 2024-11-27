package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Comment struct {
	Author    string
	Message   string
	CreatedAt int64
}

var db *gorm.DB

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", GET)
	mux.HandleFunc("POST /", POST)

	log.Print("Starting server...")
	prepareDB()
	http.ListenAndServe(":8080", mux)
}

func prepareDB() {
	var err error
	dsn := "host=localhost user=test password=123 dbname=swa port=5432 sslmode=disable TimeZone=Europe/Berlin"
	if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		panic("failed to connect database")
	}

	db.Migrator().DropTable(&Comment{})
	db.AutoMigrate(&Comment{})
}

func GET(res http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}

	var Comments []Comment
	db.
		Order("created_at desc").
		Find(&Comments)

	if err := tmpl.Execute(res, Comments); err != nil {
		panic(err)
	}
}

func POST(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	author := req.FormValue("author")
	message := req.FormValue("message")

	db.Create(&Comment{Author: author, Message: message, CreatedAt: time.Now().UnixMilli()})

	GET(res, req)
}
