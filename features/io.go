package features

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

func DemonstrateFileIO() {
	log.Info("\n--- File I/O ---")
	content := "Hello from Go!"
	err := ioutil.WriteFile("test.txt", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
	log.Info("Wrote to test.txt")

	readContent, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	log.Infof("Read from test.txt: %s", readContent)
	os.Remove("test.txt")
	log.Info("Deleted test.txt")
}

func DemonstrateHTTPClient() {
	log.Info("\n--- HTTP Client ---")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Error("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error:", err)
		return
	}
	log.Infof("HTTP Response: %s", body)
}