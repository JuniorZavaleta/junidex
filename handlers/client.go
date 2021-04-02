package handlers

import (
	"html/template"
	"junidex/helpers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ClientHandler struct {
	// can add services
}

func NewClientHandler(r *mux.Router) {
	handler := &ClientHandler{}

	r.HandleFunc("/app/starter", handler.starterView).Methods("GET")
	r.HandleFunc("/app/team", handler.teamView).Methods("GET")
}

func (h *ClientHandler) starterView(w http.ResponseWriter, r *http.Request) {
	team := &PokemonTeam{}
	t, err := template.ParseFiles(helpers.GetTemplateFilepath("client", "starter.html"))

	if err != nil {
		log.Panic(err)
	}

	t.Execute(w, team)
}

func (h *ClientHandler) teamView(w http.ResponseWriter, r *http.Request) {
	team := &PokemonTeam{}
	// TODO: Should load team in database
	t, err := template.ParseFiles(helpers.GetTemplateFilepath("client", "team.html"))

	if err != nil {
		log.Panic(err)
	}

	t.Execute(w, team)
}
