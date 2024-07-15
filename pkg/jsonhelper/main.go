package jsonhelper

type SimpleJson map[string]interface{}

func Cli(args ...string) {
	filepath, command := args[0], args[1]
	switch command {
	case "get":
		JsonGet(filepath, args[2:]...)
	case "bump":
		JsonBump(filepath, args[2:]...)
	case "set":
		JsonSet(filepath, args[2:]...)
	}
}
