package main

import (
	"log"
	"os"
	"testing"
)

const (
	testJsonIn = `{
  "project": "test",
  "version": "1.0.0",
  "commit": "abcdef1234",
  "date": "2020-09-30T00:00:00Z",
  "build_id": 0
}`
	testJsonOut = `{
  "build_id": 2,
  "commit": "abcdef1234",
  "date": "2020-09-30T00:00:00Z",
  "project": "test",
  "version": "1.1.1"
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

	functionDoIt(file.Name(), "bump", "patch", "version")
	functionDoIt(file.Name(), "bump", "int", "build_id")
	functionDoIt(file.Name(), "bump", "minor", "version")
	functionDoIt(file.Name(), "bump", "int", "build_id")
	//functionDoIt(file.Name(), "bump", "date", "date")
	//fmt.Printf("%+v\n", result)
	byteValue, _ := os.ReadFile(file.Name())

	if string(byteValue) != testJsonOut {
		t.Errorf("Got wrong content: \n %v\n", string(byteValue))
	}

}
