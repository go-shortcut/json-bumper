package main

import (
	"encoding/json"
	"fmt"
	semver "github.com/blang/semver/v4"
	"log"
	"os"
	"time"
)

const (
	banner = `
how to use...
`
)

func main() {
	if len(os.Args) < 5 {
		log.Fatal(banner)
	}
	fmt.Println(os.Args[1:5])
	functionDoIt(os.Args[1], os.Args[2], os.Args[3], os.Args[4])
}
func functionDoIt(filepath, command, option, key string) {
	var result map[string]interface{}
	switch command {
	case "get":
		byteValue, _ := os.ReadFile(filepath)
		json.Unmarshal(byteValue, &result)
		value, ok := result[key]
		if ok {
			fmt.Println(value)
		}
		// TBD get with different types
	case "bump":
		byteValue, _ := os.ReadFile(filepath)
		json.Unmarshal(byteValue, &result)
		value, ok := result[key]
		switch option {
		case "int":
			if !ok {
				value = 0
			}
			if v, ok := value.(float64); ok {
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
					log.Fatal(err)
				}
				if option == "patch" {
					version.Patch += 1
				} else if option == "minor" {
					version.Minor += 1

				} else if option == "major" {
					version.Major += 1

				}
				result[key] = version.FinalizeVersion()
			}
		}
		file, _ := json.MarshalIndent(result, "", "  ")
		_ = os.WriteFile(filepath, file, 0644)
	}
}
