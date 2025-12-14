package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-hello-world/features"
	"time"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	features.InitConfig()

	features.DemonstrateStructsAndInterfaces()
	features.DemonstrateJSON()

	log.Info("\n--- Goroutines and Channels ---")
	myChannel := make(chan string)
	go features.LongRunningTask(myChannel)
	result := <-myChannel
	log.Info(result)

	features.DemonstrateErrorHandling()
	features.DemonstrateDefer()
	features.DemonstrateSlices()
	features.DemonstrateMaps()
	features.DemonstratePointers()
	features.DemonstrateSelect()
	features.DemonstrateMutex()
	features.DemonstrateStringManipulation()
	features.DemonstrateFileIO()

	ctx, cancel := context.WithCancel(context.Background())
	go features.DemonstrateContext(ctx)
	time.Sleep(1 * time.Second)
	cancel()

	features.DemonstrateHTTPClient()

	time.Sleep(2 * time.Second)
}
