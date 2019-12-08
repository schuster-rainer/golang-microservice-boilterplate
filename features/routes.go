package features

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/schuster-rainer/myservice/features/todo"
)

// https://github.com/tonyalaribe/todoapi

// Routes for application
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		myCustomAuthMiddleware,
	)

	router.Route("/api/v1", featureRoutes)

	return router
}

func myCustomAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If user is loggedin, allows access.
		if true {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(403), 403)
			return
		}

		return
	})
}

// Routes all features
func featureRoutes(r chi.Router) {
	//TODO: Add package reflection or some other means to automatically import routes
	r.Mount("/todo", todo.Routes())
}
