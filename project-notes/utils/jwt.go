package utils

import (
    "net/http"
    "strings"
    "github.com/dgrijalva/jwt-go"
    "context"
)

func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }

        tokenString := strings.Split(authHeader, " ")[1]
        token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("secret"), nil
        })

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            ctx := context.WithValue(r.Context(), "userID", uint(claims["user_id"].(float64)))
            next.ServeHTTP(w, r.WithContext(ctx))
        } else {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
        }
    })
}
