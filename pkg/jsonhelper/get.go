package jsonhelper

import (
	"encoding/json"
	"github.com/blang/semver/v4"
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
		value, ok := result[key]
		if !ok {
			PrintToStdErr(key + " key not found.")
		}
		if v, ok := value.(string); ok {
			version, err := semver.Make(v)
			if err != nil {
				PrintToStdErrFatal(err)
				return
			}
			if option == "patch" {
				PrintToStdOut(version.Patch)
			} else if option == "minor" {
				PrintToStdOut(version.Minor)

			} else {
				// if option == "major"
				PrintToStdOut(version.Major)
			}
		}
	}
}
