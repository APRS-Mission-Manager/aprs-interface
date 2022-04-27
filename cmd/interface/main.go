package main

import (
	"os"

	"github.com/APRS-Mission-Manager/aprs-interface/internal/aprs"
	"github.com/APRS-Mission-Manager/aprs-interface/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func setup() config.Config {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Set file name of the config
	// TODO: Make this a command line argument
	viper.SetConfigName("debug")
	// Set path to find config file
	viper.AddConfigPath("../../config")
	// Allow viper to read environment variables
	viper.AutomaticEnv()
	// Set config file type
	viper.SetConfigType("yaml")

	var cfg config.Config

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("Unable to load config file!")
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("Unable to unmarshall config data!")
	}

	cfg.APRS.Username = viper.GetString("APRS_Username")
	cfg.APRS.Password = viper.GetString("APRS_Password")

	// Validate that the cfg file was correctly unmarshalled
	log.Debug().Str("aprs server", cfg.APRS.Server).Int("aprs port", cfg.APRS.Port).Str("db name", cfg.Amazon.DBAPRSLog.Name).Msg("Configuration file loaded.")

	return cfg
}

func main() {
	appConfig := setup()

	aprsHook := aprs.CreateHook(appConfig)
	go aprsHook.Subscribe()

	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(w, "Hello, world!\n")
	// }

	// http.HandleFunc("/hello", helloHandler)
	// log.Info().Msg("Listening for requests at http://localhost:8080/hello")

	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal().AnErr("error", err)
	// }
	select {}
}
