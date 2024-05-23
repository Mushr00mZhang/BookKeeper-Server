package main

import (
	"financialrecord-backend/controllers"
	"financialrecord-backend/models/database"
	"log"
	"net/http"

	// "gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func main() {

	// os.Remove("FinancialRecord.db")
	log.Println("数据库连接中。")
	db, err := gorm.Open(sqlite.Open("FinancialRecord.db"), &gorm.Config{})
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
	controllers.Init(r)
	r.PathPrefix("").Handler(http.FileServer(http.Dir("./dist")))
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
