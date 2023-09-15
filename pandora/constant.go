package pandora

import "os"

const (
	pandoraPortEnvKey  = "PANDORA_PORT_ENV"
	pandoraDefaultPort = "50051"
)

var pandoraServicePort string

func getPort() {
	pandoraServicePort = os.Getenv("PORT")
	if pandoraServicePort == "" {
		pandoraServicePort = pandoraDefaultPort
	}
}
