package todo

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Todo type
type Todo struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Routes for Todo type
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoID}", Get)
	router.Delete("/{todoID}", Delete)
	router.Post("/", Create)
	router.Get("/", Index)
	return router
}

// Get a todo
func Get(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todos := Todo{
		Slug:  todoID,
		Title: "Hello world",
		Body:  "Heloo world from planet earth",
	}
	render.JSON(w, r, todos) // A chi router helper for serializing and returning json
}

// Delete a todo
func Delete(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted TODO successfully"
	render.JSON(w, r, response) // Return some demo response
}

// Create todo
func Create(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created TODO successfully"
	render.JSON(w, r, response) // Return some demo response
}

// Index all todos
func Index(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{
			Slug:  "slug",
			Title: "Hello world",
			Body:  "Heloo world from planet earth",
		},
	}
	render.JSON(w, r, todos) // A chi router helper for serializing and returning json
}
