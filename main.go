package main

import (
	"log"

	"github.com/pedropezzuol/api_test/config"
	"github.com/pedropezzuol/api_test/controller/routes"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}

	routes.InitializeRouter()
}
