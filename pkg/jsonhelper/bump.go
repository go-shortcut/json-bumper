package jsonhelper

import (
	"encoding/json"
	"github.com/blang/semver/v4"
	"os"
	"time"
)

func JsonBump(fpath string, args ...string) {
	var result SimpleJson
	option, key := args[0], args[1]
	byteValue, _ := os.ReadFile(fpath)
	_ = json.Unmarshal(byteValue, &result)
	value, ok := result[key]

	switch option {
	case "int":
		if !ok {
			value = float64(0)
		}
		if v, ok2 := value.(float64); ok2 {
			result[key] = v + 1
		}
	case "date":
		now := time.Now().UTC()
		result[key] = now.Format(time.RFC3339)
	case "patch", "minor", "major":
		if !ok {
			value = ""
		}
		if v, ok := value.(string); ok {
			version, err := semver.Make(v)
			if err != nil {
				PrintToStdErrFatal(err)
			}
			if option == "patch" {
				version.Patch += 1
			} else if option == "minor" {
				version.Minor += 1

			} else {
				// if option == "major"
				version.Major += 1
			}
			result[key] = version.FinalizeVersion()
		}
	}

	WriteFormattedJsonToFile(fpath, result)

}
