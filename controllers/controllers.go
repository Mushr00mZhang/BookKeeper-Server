package controllers

import (
	"financialrecord-backend/controllers/outlay"
	"financialrecord-backend/controllers/outlaycat"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()
	outlay.Init(api)
	outlaycat.Init(api)
}
