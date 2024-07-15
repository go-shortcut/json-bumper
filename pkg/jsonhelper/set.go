package jsonhelper

import (
	"encoding/json"
	"os"
)

func JsonSet(fpath string, args ...string) {
	// string commit git_hash
	var result SimpleJson
	_, key, newValue := args[0], args[1], args[2]
	byteValue, _ := os.ReadFile(fpath)
	_ = json.Unmarshal(byteValue, &result)
	result[key] = newValue
	_ = WriteFormattedJsonToFile(fpath, result)
}
