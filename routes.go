package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *Main) routes() *chi.Mux {
	//middleware must come before any routes

	//add routes
	//pages
	a.get("/", a.Handlers.Home)
	a.get("/download/{book}", a.Handlers.DownloadFile)

	//api

	//static routes
	fileserver := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileserver))
	return a.App.Routes
}
