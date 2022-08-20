package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8082"

type Config struct{}

func main() {
	app := Config{}

	log.Println("Starting broker service on port", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
