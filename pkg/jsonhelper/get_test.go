package jsonhelper

import (
	"log"
	"os"
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

	versionP := captureStdout(func() {
		Cli(file.Name(), "get", "patch", "version")
	})
	if versionP != "9" {
		t.Errorf("wrong patch version: '%s'", versionP)
	}
	versionMi := captureStdout(func() {
		Cli(file.Name(), "get", "minor", "version")
	})
	if versionMi != "0" {
		t.Errorf("wrong minor version: '%s'", versionMi)
	}
	versionMa := captureStdout(func() {
		Cli(file.Name(), "get", "major", "version")
	})
	if versionMa != "1" {
		t.Errorf("wrong major version: '%s'", versionMa)
	}
	emptyString := captureStdout(func() {
		Cli(file.Name(), "get", "int", "build_id")
	})
	if emptyString != "" {
		t.Errorf("expected empty string")
	}

	verString := captureStdout(func() {
		Cli(file.Name(), "get", "string", "version")
	})

	if verString != "1.0.9" {
		t.Errorf("Got wrong version: '%v'\n", verString)
	}

	// run after test to fill coverage gaps
	emptyString = captureStdout(func() {
		Cli(file.Name(), "get", "patch", "project")
	})
	if emptyString != "" {
		t.Errorf("expected empty string")
	}

	emptyString = captureStdout(func() {
		Cli(file.Name(), "get", "major", "not_existing_key")
	})
	if emptyString != "" {
		t.Errorf("expected empty string")
	}
}
