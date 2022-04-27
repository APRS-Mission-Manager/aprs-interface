package aprs

import (
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/APRS-Mission-Manager/aprs-interface/internal/config"
	"github.com/rs/zerolog/log"
)

type Hook struct {
	conn         net.Conn
	appConfig    config.Config
	packetLogger APRSPacketLogger
}

type APRSPacketLogger interface {
	LogAPRSPacket(callsign string, timestamp int64, packet string)
}

func CreateHook(appConfig config.Config, packetLogger APRSPacketLogger) Hook {
	hook := Hook{appConfig: appConfig, packetLogger: packetLogger}
	return hook
}

func (hook Hook) Subscribe() {
	log.Info().Msg("[APRS Hook] Initializing APRS Hook.")

	// Create TCP Connection to APRS-IS server
	serverAddress := hook.appConfig.APRS.Server + ":" + strconv.Itoa(hook.appConfig.APRS.Port)
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("[APRS Hook] Failed while connecting to " + serverAddress)
	}
	hook.conn = conn

	// Connection needs to be followed by login, otherwise the server will close the connection
	hook.login()
	replyBuffer := make([]byte, 256)
	for {
		_, err := conn.Read(replyBuffer)
		if err != nil {
			log.Fatal().AnErr("error", err).Msg("[APRS Hook] Failed while receiving message from server.")
		}

		message := strings.Split(string(replyBuffer), "\r\n")[0]

		// Non-APRS packets are prefixed by '#'
		if message[0] == '#' {
			hook.handleServerUpdate(message)
		} else {
			hook.handleAPRSPacket(message)
		}
	}
}

func (hook Hook) login() {
	filterMsg := " filter"
	for i := range hook.appConfig.APRS.CallsignPatterns {
		filterMsg = filterMsg + " " + hook.appConfig.APRS.CallsignPatterns[i]
	}
	loginMsg := "user " + hook.appConfig.APRS.Username + " pass " + hook.appConfig.APRS.Password + filterMsg + "\r\n"

	_, err := hook.conn.Write([]byte(loginMsg))
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("[APRS Hook] Failed while sending login message to server.")
	}
}

func (hook Hook) handleAPRSPacket(packet string) {
	log.Debug().Str("packet", packet).Msg("[APRS Hook]")

	// Packets are split <callsign> '>' <rest of packet>
	callsign := strings.Split(packet, ">")[0]
	// Timestamp is an internal measurement of when the packet was processed
	timestamp := time.Now().UnixMilli()

	hook.packetLogger.LogAPRSPacket(callsign, timestamp, packet)
}

func (hook Hook) handleServerUpdate(update string) {
	log.Debug().Str("update", update).Msg("[APRS Hook]")
}
