package main

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-features-showcase/features"
	"go-features-showcase/server"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})

	features.InitConfig()

	level, err := logrus.ParseLevel(viper.GetString("log_level"))
	if err != nil {
		log.Fatalf("failed to parse log level: %v", err)
	}
	log.SetLevel(level)

	s := server.NewServer()
	s.Run()
}
