package main

import (
	"base/data"
	"base/handlers"
	"base/middleware"

	"github.com/m-goku/rkt"
)

type Main struct {
	App        *rkt.RKT
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	rocket := initApplication()
	rocket.App.ListenAndServe()
}
