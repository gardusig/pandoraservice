package internal

import "os"

const (
	pandoraPortEnvKey = "PANDORA_GRPC_PORT"
	pandoraHostEnvKey = "PANDORA_GRPC_HOST"

	pandoraDefaultPort = "50051"
	pandoraDefaultHost = "localhost"
)

var (
	PandoraServicePort string
	PandoraServerHost  string
)

var (
	Equal   = "="
	Less    = "<"
	Greater = ">"
)

const (
	LevelMinThreshold uint32 = 0
	LevelMaxThreshold uint32 = 1000

	GuessMinThreshold int64 = -4000000000000000000
	GuessMaxThreshold int64 = +4000000000000000000
)

func init() {
	PandoraServicePort = os.Getenv(pandoraPortEnvKey)
	if PandoraServicePort == "" {
		PandoraServicePort = pandoraDefaultPort
	}
	PandoraServerHost = os.Getenv(pandoraHostEnvKey)
	if PandoraServerHost == "" {
		PandoraServerHost = pandoraDefaultHost
	}
}
