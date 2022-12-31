package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		//AllowOriginFunc:    nil,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		//OptionsPassthrough: false,
		//Debug:              false,
	}))

	mux.Post("/users/login", app.Login)
	mux.Post("/users/logout", app.Logout)

	mux.Post("/books", app.AllBooks)
	mux.Get("/books", app.AllBooks)

	mux.Get("/books/{slug}", app.GetOneBook)

	mux.Post("/validate-token", app.ValidateToken)

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.AuthTokenMiddleware)

		mux.Post("/users", app.AllUsers)

		mux.Post("/users/save", app.EditUser)
		mux.Post("/users/get/{id}", app.GetUser)

		mux.Post("/users/delete", app.DeleteUser)
		mux.Post("/log-user-out/{id}", app.LogUserOutAndSetInactive)

		mux.Post("/authors/all", app.AuthorsAll)
		mux.Post("/books/save", app.EditBook)

		mux.Post("/books/delete", app.DeleteBook)
		mux.Post("/books/{id}", app.BookById)

	})

	//static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
