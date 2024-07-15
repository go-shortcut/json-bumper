package jsonhelper

import (
	"encoding/json"
	"os"
)

func JsonGet(fpath string, args ...string) {
	var result SimpleJson

	option, key := args[0], args[1]
	byteValue, _ := os.ReadFile(fpath)
	_ = json.Unmarshal(byteValue, &result)

	switch option {
	case "string":
		value, ok := result[key]
		if ok {
			PrintToStdOut(value)
		}
	case "int":
		PrintToStdErr("'int' does not make sense here.")
	case "patch", "minor", "major":
		PrintToStdErr("'path,minor,major' are not implemented yet.")
	}
}
