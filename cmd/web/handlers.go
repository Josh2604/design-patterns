package main

import (
	"fmt"
	"net/http"

	"github.com/Josh2604/design-patterns/pets"
	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	// dog := pets.NewPet("dog")
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))

}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	// dog := pets.NewPet("cat")
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))

}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}
