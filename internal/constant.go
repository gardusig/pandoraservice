package internal

import "os"

const (
	pandoraPortEnvKey  = "PANDORA_PORT_ENV"
	pandoraDefaultPort = "50051"
)

const (
	MinThreshold = -4000000000000000000
	MaxThreshold = +4000000000000000000
)

var (
	Equal   = "="
	Less    = "<"
	Greater = ">"
)

var EncryptedMessage = "important encrypted message"

var PandoraServicePort string

func init() {
	PandoraServicePort = os.Getenv(pandoraPortEnvKey)
	if PandoraServicePort == "" {
		PandoraServicePort = pandoraDefaultPort
	}
}