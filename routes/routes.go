package routes

import (
    "github.com/gorilla/mux"
    "gorm.io/gorm"
    "your-go-app/views"
)

func RegisterRoutes(r *mux.Router, db *gorm.DB) {
    r.HandleFunc("/users", views.GetUsers(db)).Methods("GET")
}
