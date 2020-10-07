package main

import (
	"bytes"
	"github.com/rs/zerolog/log"
	"os"
	"os/exec"
	"time"
)

const binName = "scripts/run.sh"

func main() {
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)
	cmd := &exec.Cmd{
		Path:   binName,
		Stdout: bufOut,
		Stderr: bufErr,
	}
	cmd.Env = os.Environ()

	dir, _:= os.Getwd()

	log.Info().Msgf("pwd {%s}", dir)

	go printOut(bufOut)
	go printOut(bufErr)

	err := cmd.Run()
	if err != nil {
		log.Error().Err(err).Msg("failed run")
	}

	log.Info().Msgf("bufOut: %s", bufOut.String())
	log.Info().Msgf("bufErr: %s", bufErr.String())

}

func printOut(out *bytes.Buffer)  {
	lastLen := 0
	for {
		len := out.Len()
		if len > lastLen {
			bytes := out.Bytes()
			part := bytes[lastLen:len]
			log.Info().Msgf("part {%s}", part)
		}
		lastLen = len
		time.Sleep(time.Millisecond * 100)
	}
}
