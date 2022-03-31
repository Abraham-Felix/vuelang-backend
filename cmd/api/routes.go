package main

// here we import the packages we will use similiar to node modules
import (
  "net/http"
  "github.com/go-chi/chi/v5"
  "github.com/go-chi/chi/v5/middleware"
)

// function bellow has the application as receiver
// the function name is routes and returns http handler
// inside there is a router function called mux using chi version 5 package
func (app *application) routes() http.Handler {
    mux := chi.NewRouter()
    mux.Use(middleware.Recoverer)

    //this for local
    mux.Get("/users/login", app.Login)
    //this to post to front end 
    mux.Post("/users/login", app.Login)

    return mux
}
