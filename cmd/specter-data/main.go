package main

import (
	"github.com/namsral/flag"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"go.smantic.dev/access-viz/internal/dataServer"
)

var logLevel string

func init() {
	flag.StringVar(&logLevel, "loglvl", "Warn", "Log level")
}

func main() {
	flag.Parse()
	switch {
	case logLevel == "Debug":
		log.SetLevel(logrus.DebugLevel)
	case logLevel == "Info":
		log.SetLevel(logrus.InfoLevel)
	case logLevel == "Warn":
		log.SetLevel(logrus.WarnLevel)
	case logLevel == "Error":
		log.SetLevel(logrus.ErrorLevel)
	case logLevel == "Fatal":
		log.SetLevel(logrus.FatalLevel)
	}

	dataServer.Start()
}
