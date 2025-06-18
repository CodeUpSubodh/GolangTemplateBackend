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
