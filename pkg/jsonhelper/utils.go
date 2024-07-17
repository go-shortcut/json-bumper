package jsonhelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func PrintToStdErr(msg string) {
	println(msg)
}

func PrintToStdErrFatal(anything interface{}) {
	fmt.Fprintln(os.Stderr, anything)
}

func PrintToStdOut(anything interface{}) {
	fmt.Println(anything)
}

func WriteFormattedJsonToFile(fpath string, sjson SimpleJson) (err error) {

	content, _ := json.MarshalIndent(sjson, "", "  ")
	err = os.WriteFile(fpath, content, 0644)
	return err
}

// not thread safe
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return strings.TrimSuffix(buf.String(), "\n")
}
