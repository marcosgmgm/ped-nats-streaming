package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {

	nc, err := stan.Connect("dennis-nats-stan-logs", "write-1")
	if err != nil {
		log.Fatal().Err(err).Msg("Fail to connect to nats.")
	}
	i := 0
	for {
		err := nc.Publish("test", []byte(fmt.Sprintf("test-%d", i)))
		if err != nil {
			log.Error().Err(err).Msg("failed to publish nats")
		}
		log.Info().Msg(fmt.Sprintf(fmt.Sprintf("test-%d", i)))
		time.Sleep(time.Second * 1)
		i++
	}


}
