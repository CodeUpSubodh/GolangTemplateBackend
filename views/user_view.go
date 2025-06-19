package views

import (
	"encoding/json"
	"net/http"
	"your-go-app/models"

	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []models.User
		db.Find(&users)
		json.NewEncoder(w).Encode(users)
	}
}

func CreateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		// Decode the request body into user struct
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Basic validation (you can enhance this)
		if user.Name == "" || user.Email == "" {
			http.Error(w, "Name and Email are required fields", http.StatusBadRequest)
			return
		}

		// Create user in the database
		if err := db.Create(&user).Error; err != nil {
			http.Error(w, "Could not create user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond with the created user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
