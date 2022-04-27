package main

import (
	"io"
	"net/http"
	"os"

	"github.com/APRS-Mission-Manager/aprs-interface/internal/aprs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	go aprs.InitializeHook()

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Info().Msg("Listening for requests at http://localhost:8080/hello")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal().AnErr("error", err)
	}
}
