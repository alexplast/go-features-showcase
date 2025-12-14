package features

import (
	"context"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func DemonstrateFileIO() {
	log.Info("\n--- File I/O ---")
	content := "Hello from Go!"
	err := os.WriteFile("test.txt", []byte(content), 0o644)
	if err != nil {
		panic(err)
	}
	log.Info("Wrote to test.txt")

	readContent, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	log.Infof("Read from test.txt: %s", readContent)
	if err := os.Remove("test.txt"); err != nil {
		log.Error("Error:", err)
	}
	log.Info("Deleted test.txt")
}

func DemonstrateHTTPClient() {
	log.Info("\n--- HTTP Client ---")
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		log.Error("Error:", err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Error:", err)
		return
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Error("Error:", closeErr)
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error:", err)
		return
	}
	log.Infof("HTTP Response: %s", body)
}
