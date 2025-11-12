package main

import (
	"net/http"
)

func (a *Main) get(s string, h http.HandlerFunc) {
	a.App.Routes.Get(s, h)
}

func (a *Main) post(s string, h http.HandlerFunc) {
	a.App.Routes.Post(s, h)
}

func (a *Main) use(m ...func(http.Handler) http.Handler) {
	a.App.Routes.Use(m...)
}
