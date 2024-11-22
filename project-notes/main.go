package main

import (
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "project-notes/models"
    "project-notes/routes"
)

var DB *gorm.DB

func main() {
    var err error
    DB, err = gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer DB.Close()

    // Migrasi tabel ke database
    DB.AutoMigrate(&models.User{}, &models.Note{})

    // Atur routing dan jalankan server
    router := routes.SetupRoutes()
    http.ListenAndServe(":8000", router)
}
