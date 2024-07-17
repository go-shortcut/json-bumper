package jsonhelper

import (
	"log"
	"os"
	"testing"
)

const (
	testJsonIn = `{
  "project": "test",
  "version": "1.0.9",
  "commit": "abcdef1234",
  "date": "2020-09-30T00:00:00Z"
}`
	testJsonOut = `{
  "build_id": 2,
  "commit": "567890abcd",
  "date": "2020-09-30T00:00:00Z",
  "project": "test",
  "version": "2.1.10"
}`
)

func TestName(t *testing.T) {

	file, err := os.CreateTemp(os.TempDir(), "json-bumper-")
	if err != nil {
		log.Fatal(err)
	}
	//defer os.Remove(file.Name())
	file.WriteString(testJsonIn)
	file.Close()

	Cli(file.Name(), "bump", "patch", "version")
	Cli(file.Name(), "bump", "int", "build_id")
	Cli(file.Name(), "bump", "minor", "version")
	Cli(file.Name(), "bump", "int", "build_id")
	Cli(file.Name(), "bump", "major", "version")
	Cli(file.Name(), "set", "string", "commit", "567890abcd")

	byteValue, _ := os.ReadFile(file.Name())

	if string(byteValue) != testJsonOut {
		t.Errorf("Got wrong content: \n %v\n", string(byteValue))
	}

	// run after test to fill coverage gaps
	Cli(file.Name(), "bump", "date", "date")
	Cli(file.Name(), "bump", "major", "not_existing_key")

}
