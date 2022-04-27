package aprs

import (
	"net"
	"strings"

	"github.com/rs/zerolog/log"
)

func InitializeHook() {
	log.Info().Msg("Initializing APRS Hook.")

	conn, err := net.Dial("tcp", "noam.aprs2.net:14580")
	checkErr(err)

	defer conn.Close()

	loginMsg := "user kn4cdd pass -1 filter r/33/-97/200\r\n"

	_, err = conn.Write([]byte(loginMsg))
	checkErr(err)

	replyBuffer := make([]byte, 128)
	replyBuffer[127] = 'A'

	for {
		_, err = conn.Read(replyBuffer)
		checkErr(err)

		packet := strings.Split(string(replyBuffer), "\r\n")[0]

		log.Info().Msg(packet)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal().AnErr("error", err)
	}
}
