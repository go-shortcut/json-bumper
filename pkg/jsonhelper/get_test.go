package jsonhelper

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {

	file, err := os.CreateTemp(os.TempDir(), "json-bumper-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	file.WriteString(testJsonIn)
	file.Close()

	// Cli(file.Name(), "get", "patch", "version")
	// Cli(file.Name(), "get", "minor", "version")
	emptyString := captureStdout(func() {
		Cli(file.Name(), "get", "major", "version")
	})
	if emptyString != "" {
		t.Errorf("expected empty string, bacause is not implemented")
	}
	emptyString = captureStdout(func() {
		Cli(file.Name(), "get", "int", "build_id")
	})
	if emptyString != "" {
		t.Errorf("expected empty string")
	}

	verString := captureStdout(func() {
		Cli(file.Name(), "get", "string", "version")
	})

	if strings.TrimSuffix(verString, "\n") != "1.0.0" {
		t.Errorf("Got wrong version: '%v'\n", verString)
	}

}
