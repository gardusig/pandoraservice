package internal

import "os"

const (
	pandoraPortEnvKey  = "PANDORA_PORT_ENV"
	pandoraDefaultPort = "50051"
)

const (
	LevelMinThreshold uint32 = 0
	LevelMaxThreshold uint32 = 1000

	GuessMinThreshold int64 = -4000000000000000000
	GuessMaxThreshold int64 = +4000000000000000000
)

var (
	Equal   = "="
	Less    = "<"
	Greater = ">"
)

var PandoraServicePort string

func init() {
	PandoraServicePort = os.Getenv(pandoraPortEnvKey)
	if PandoraServicePort == "" {
		PandoraServicePort = pandoraDefaultPort
	}
}
