package controllers

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "project-notes/models"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
    var note models.Note
    json.NewDecoder(r.Body).Decode(&note)
    
    // Ambil userID dari JWT token
    userID := r.Context().Value("userID").(uint)
    note.UserID = userID

    // Simpan note ke database
    models.DB.Create(&note)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(note)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(uint)
    var notes []models.Note

    // Ambil notes milik user yang sedang login
    models.DB.Where("user_id = ?", userID).Find(&notes)
    json.NewEncoder(w).Encode(notes)
}
