package handlers

import (
	"base/data"
	"fmt"
	"net/http"

	"github.com/m-goku/rkt"
)

type Handlers struct {
	App    *rkt.RKT
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering", err)
	}
}



// JSON is the handler to demonstrate json responses
// this tests how json response is sent
func (h *Handlers) JSON(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID      int64    `json:"id"`
		Name    string   `json:"name"`
		Hobbies []string `json:"hobbies"`
	}

	payload.ID = 10
	payload.Name = "Jack Jones"
	payload.Hobbies = []string{"karate", "tennis", "programming"}

	err := h.App.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

// DownloadFile is the handler to demonstrate file download reponses
func (h *Handlers) DownloadFile(w http.ResponseWriter, r *http.Request) {
	book := r.PathValue("book")
	h.App.DownloadFile(w, r, "./public/books", fmt.Sprintf("%s.pdf", book))
}
