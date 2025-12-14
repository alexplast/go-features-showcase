package main

import (
	log "github.com/sirupsen/logrus"
	"go-features-showcase/features"
	"go-features-showcase/server"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	features.InitConfig()

	s := server.NewServer()
	s.Run()
}
