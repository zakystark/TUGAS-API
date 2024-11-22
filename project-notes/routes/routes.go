package routes

import (
    "github.com/gorilla/mux"
    "project-notes/controllers"
    "project-notes/utils"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Routes untuk register dan login
    router.HandleFunc("/register", controllers.Register).Methods("POST")
    router.HandleFunc("/login", controllers.Login).Methods("POST")

    // Routes untuk notes (dilindungi JWT)
    notesRouter := router.PathPrefix("/notes").Subrouter()
    notesRouter.Use(utils.JWTMiddleware)
    notesRouter.HandleFunc("", controllers.CreateNote).Methods("POST")
    notesRouter.HandleFunc("", controllers.GetNotes).Methods("GET")

    return router
}
