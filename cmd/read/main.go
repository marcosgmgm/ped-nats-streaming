package main

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {

	natc, err := nats.Connect(
		"nats://localhost:4222",
		nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
			log.Err(err).Msg("nats client disconnected")
		}),
		nats.ReconnectHandler(func(_ *nats.Conn) {
			log.Info().Msg("nats client reconnected")
		}),
		nats.ClosedHandler(func(_ *nats.Conn) {
			log.Info().Msg("nats client closed")
		}),
	)
	nc, err := stan.Connect("dennis-nats-stan-logs", "read-1", stan.NatsConn(natc))
	if err != nil {
		log.Fatal().Err(err).Msg("Fail to connect to nats.")
	}

	hadler := printData

	_, err = nc.QueueSubscribe("test", "read-01", hadler, stan.StartAt(4))
	if err != nil {
		log.Error().Err(err).Msg("error chan subscribe")
	}

	time.Sleep(time.Hour)

}

func printData(m *stan.Msg) {
	log.Info().Msg(string(m.Data))
}
