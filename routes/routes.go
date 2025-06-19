package routes

import (
	"your-go-app/views"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/users", views.GetUsers(db)).Methods("GET")
	r.HandleFunc("/user", views.GetUsers(db)).Methods("POST")

}
