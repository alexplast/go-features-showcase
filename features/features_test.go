package features

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestDemonstrateStructsAndInterfaces(t *testing.T) {
	DemonstrateStructsAndInterfaces()
}

func TestDemonstrateJSON(t *testing.T) {
	DemonstrateJSON()
}

func TestDemonstrateSlices(t *testing.T) {
	DemonstrateSlices()
}

func TestDemonstrateMaps(t *testing.T) {
	DemonstrateMaps()
}

func TestDemonstrateFileIO(t *testing.T) {
	DemonstrateFileIO()
}

func TestDemonstrateHTTPClient(t *testing.T) {
	DemonstrateHTTPClient()
}

func TestDemonstrateSelect(t *testing.T) {
	DemonstrateSelect()
}

func TestDemonstrateMutex(t *testing.T) {
	DemonstrateMutex()
}

func TestDemonstrateContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	DemonstrateContext(ctx)
}

func TestDemonstrateErrorHandling(t *testing.T) {
	DemonstrateErrorHandling()
}

func TestDemonstrateDefer(t *testing.T) {
	DemonstrateDefer()
}

func TestInitConfig(t *testing.T) {
	t.Cleanup(func() {
		os.Remove("config.yml")
	})
	// Create a dummy config file
	f, err := os.Create("config.yml")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString("greeting: hello")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	InitConfig()
}
