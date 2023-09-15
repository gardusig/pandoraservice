package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {

}
