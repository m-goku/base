package main

import (
	"base/data"
	"base/handlers"
	"base/middleware"
	"log"
	"os"

	"github.com/m-goku/rkt"
)

func initApplication() *Main {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init rkt
	rkt := &rkt.RKT{}
	err = rkt.New(path)
	if err != nil {
		log.Fatal(err)
	}

	rkt.AppName = "base"

	myMiddleware := &middleware.Middleware{
		App: rkt,
	}

	myHandlers := &handlers.Handlers{
		App: rkt,
	}

	base := &Main{
		App:        rkt,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}
	base.App.Routes = base.routes()
	base.Models = data.NewModel(base.App.DB.Pool)
	myHandlers.Models = base.Models
	base.Middleware.Models = base.Models

	return base
}
