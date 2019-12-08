package main

import (
	"log"

	"github.com/schuster-rainer/myservice/cmd"
)

func main() {
	// config structure from
	// https://www.netlify.com/blog/2016/09/06/creating-a-microservice-boilerplate-in-go/
	// https://github.com/rybit/seltzer/tree/8712539a136334b604851b8cc046dbe89b11115b

	rootCmd := cmd.Server()
	rootCmd.AddCommand(cmd.PrintRoutes())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
