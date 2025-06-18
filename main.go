package main

import (
    "log"
    "net/http"
    "your-go-app/config"
    "your-go-app/models"
    "your-go-app/routes"

    "github.com/gorilla/mux"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(postgres.Open(config.DB_URL), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    db.AutoMigrate(&models.User{})

    r := mux.NewRouter()
    routes.RegisterRoutes(r, db)

    http.ListenAndServe(":8080", r)
}
