package main

import (
	"bookkeeper-backend/controllers"
	"bookkeeper-backend/models/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {

	// os.Remove("BookKeeper.db")
	log.Println("数据库连接中。")
	db, err := gorm.Open(sqlite.Open("BookKeeper.db"), &gorm.Config{})
	if err != nil {
		log.Printf("数据库连接失败。%s\n", err)
		log.Panic(1)
	}
	log.Printf("数据库已连接。\n")
	err = database.Init(db)
	if err != nil {
		log.Printf("数据库初始化失败。%s\n", err)
		log.Panic(2)
	}

	r := mux.NewRouter()
	sub := r.PathPrefix("").Subrouter()
	sub.Handle("/", http.FileServer(http.Dir("./dist/")))
	controllers.Init(sub)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
