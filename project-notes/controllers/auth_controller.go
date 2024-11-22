package controllers

import (
    "net/http"
    "encoding/json"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
    "project-notes/models"
    "project-notes/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    // Hash password
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    user.Password = string(hashedPassword)

    // Simpan user ke database
    models.DB.Create(&user)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    var input models.User
    json.NewDecoder(r.Body).Decode(&input)

    models.DB.Where("email = ?", input.Email).First(&user)

    // Periksa password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
        return
    }

    // Buat token JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, _ := token.SignedString([]byte("secret"))

    json.NewEncoder(w).Encode(map[string]string{
        "token": tokenString,
    })
}
