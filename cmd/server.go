package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/schuster-rainer/myservice/conf"
	"github.com/schuster-rainer/myservice/features"
	"github.com/spf13/cobra"
)

var serverCmd = cobra.Command{
	Use: "Run the microservice",
	Run: run,
}

// Server will setup and return the root command
func Server() *cobra.Command {
	serverCmd.PersistentFlags().StringP("config", "c", "", "the config file to use")
	serverCmd.Flags().IntP("port", "p", 4000, "the port to use")

	return &serverCmd
}

func run(cmd *cobra.Command, args []string) {
	config, err := conf.LoadConfig(cmd)
	if err != nil {
		log.Fatal("Failed to load config: " + err.Error())
	}

	logger, err := conf.ConfigureLogging(&config.LogConfig)
	if err != nil {
		log.Fatal("Failed to configure logging: " + err.Error())
	}

	logger.Infof("Starting with config: %+v", config)

	router := features.Routes()

	address := fmt.Sprintf("%s:%d", config.Host, config.Port)

	if config.Host == "" {
		log.Printf("Listening on http://localhost:%d", config.Port)
	} else {
		log.Printf("Listening on %s", address)
	}
	// // log.Print(pretty.Println(routes))
	logger.Fatal(http.ListenAndServe(address, router))
}
