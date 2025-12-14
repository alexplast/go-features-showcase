package main

import (
	log "github.com/sirupsen/logrus"
	"go-hello-world/features"
	"go-hello-world/server"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	features.InitConfig()

	s := server.NewServer()
	s.Run()
}