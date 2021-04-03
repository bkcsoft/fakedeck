package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = zerolog.New(os.Stdout).Level(zerolog.InfoLevel).With().Timestamp().Logger()
	d := VLCDeckNew()
	d.PowerOn()
}
