package cmd

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/schuster-rainer/myservice/features"
	"github.com/spf13/cobra"
)

func PrintRoutes() *cobra.Command {
	router := features.Routes()
	return &cobra.Command{
		Use:   "routes ",
		Short: "Print routes",
		Run: func(cmd *cobra.Command, args []string) {
			// Walk the route tree for every route
			walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
				log.Printf("%s %s\n", method, route)
				return nil
			}
			if err := chi.Walk(router, walkFunc); err != nil {
				// panic if there is an error
				log.Panicf("Logging err: %s\n", err.Error())
			}
		},
	}
}
