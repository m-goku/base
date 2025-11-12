package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *Main) routes() *chi.Mux {
	//middleware must come before any routes
	a.use(a.Middleware.CheckRemember)
	//add routes
	//pages
	a.get("/", a.Handlers.Home)
	a.get("/sign-up", a.Handlers.UserSignUpPage)
	a.get("/login", a.Handlers.UserLoginPage)
	a.get("/forgot-password", a.Handlers.ForgotPasswordPage)
	a.get("/reset-password", a.Handlers.ResetPasswordFormPage)
	//api
	a.post("/sign-up", a.Handlers.UserSignUp)
	a.post("/login", a.Handlers.UserLogin)
	a.post("/forgot-password", a.Handlers.ForgotPassword)
	a.post("/reset-password", a.Handlers.ResetPassword)
	a.get("/logout", a.Handlers.Logout)

	//static routes
	fileserver := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileserver))
	return a.App.Routes
}
